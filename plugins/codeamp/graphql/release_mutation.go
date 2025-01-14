package graphql_resolver

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/codeamp/circuit/plugins"
	"github.com/codeamp/circuit/plugins/codeamp/auth"
	db_resolver "github.com/codeamp/circuit/plugins/codeamp/db"
	"github.com/codeamp/circuit/plugins/codeamp/model"
	log "github.com/codeamp/logger"
	"github.com/codeamp/transistor"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

// Secret Resolver Mutation
type ReleaseResolverMutation struct {
	// DB
	DB *gorm.DB
	// Events
	Events chan transistor.Event
}

// CreateRelease
func (r *ReleaseResolverMutation) CreateRelease(ctx context.Context, args *struct{ Release *model.ReleaseInput }) (*ReleaseResolver, error) {
	var project model.Project
	var secrets []model.Secret
	var services []model.Service
	var projectExtensions []model.ProjectExtension
	var secretsJsonb postgres.Jsonb
	var servicesJsonb postgres.Jsonb
	var projectExtensionsJsonb postgres.Jsonb

	isRollback := false

	userID, err := auth.CheckAuth(ctx, []string{})
	if err != nil {
		return nil, err
	}

	// Check if project can create release in environment
	if r.DB.Where("environment_id = ? and project_id = ?", args.Release.EnvironmentID, args.Release.ProjectID).Find(&model.ProjectEnvironment{}).RecordNotFound() {
		return nil, errors.New("Project not allowed to create release in given environment")
	}

	if args.Release.ID == nil {
		projectSecrets := []model.Secret{}
		// get all the env vars related to this release and store
		r.DB.Where("environment_id = ? AND project_id = ? AND scope = ?", args.Release.EnvironmentID, args.Release.ProjectID, "project").Find(&projectSecrets)
		for _, secret := range projectSecrets {
			var secretValue model.SecretValue
			r.DB.Where("secret_id = ?", secret.Model.ID).Order("created_at desc").First(&secretValue)
			secret.Value = secretValue
			secrets = append(secrets, secret)
		}

		globalSecrets := []model.Secret{}
		r.DB.Where("environment_id = ? AND scope = ?", args.Release.EnvironmentID, "global").Find(&globalSecrets)
		for _, secret := range globalSecrets {
			var secretValue model.SecretValue
			r.DB.Where("secret_id = ?", secret.Model.ID).Order("created_at desc").First(&secretValue)
			secret.Value = secretValue
			secrets = append(secrets, secret)
		}

		secretsMarshaled, err := json.Marshal(secrets)
		if err != nil {
			return &ReleaseResolver{}, err
		}

		secretsJsonb = postgres.Jsonb{secretsMarshaled}

		r.DB.Where("project_id = ? and environment_id = ?", args.Release.ProjectID, args.Release.EnvironmentID).Find(&services)
		if len(services) == 0 {
			log.InfoWithFields("no services found", log.Fields{
				"project_id": args.Release.ProjectID,
			})
		}

		for i, service := range services {
			ports := []model.ServicePort{}
			r.DB.Where("service_id = ?", service.Model.ID).Find(&ports)
			services[i].Ports = ports

			deploymentStrategy := model.ServiceDeploymentStrategy{}
			r.DB.Where("service_id = ?", service.Model.ID).Find(&deploymentStrategy)
			services[i].DeploymentStrategy = deploymentStrategy

			readinessProbe := model.ServiceHealthProbe{}
			err = r.DB.Where("service_id = ? and type = ?", service.Model.ID, "readinessProbe").Find(&readinessProbe).Error
			if err != nil && !gorm.IsRecordNotFoundError(err) {
				return nil, err
			}
			readinessHeaders := []model.ServiceHealthProbeHttpHeader{}
			err = r.DB.Where("health_probe_id = ?", readinessProbe.ID).Find(&readinessHeaders).Error
			if err != nil && !gorm.IsRecordNotFoundError(err) {
				return nil, err
			}
			readinessProbe.HttpHeaders = readinessHeaders
			services[i].ReadinessProbe = readinessProbe

			livenessProbe := model.ServiceHealthProbe{}
			err = r.DB.Where("service_id = ? and type = ?", service.Model.ID, "livenessProbe").Find(&livenessProbe).Error
			if err != nil && !gorm.IsRecordNotFoundError(err) {
				return nil, err
			}
			livenessHeaders := []model.ServiceHealthProbeHttpHeader{}
			err = r.DB.Where("health_probe_id = ?", livenessProbe.ID).Find(&livenessHeaders).Error
			if err != nil && !gorm.IsRecordNotFoundError(err) {
				return nil, err
			}
			livenessProbe.HttpHeaders = livenessHeaders
			services[i].LivenessProbe = livenessProbe
		}

		servicesMarshaled, err := json.Marshal(services)
		if err != nil {
			return &ReleaseResolver{}, err
		}

		servicesJsonb = postgres.Jsonb{servicesMarshaled}
		// check if any project extensions that are not 'once' exists
		r.DB.Where("project_id = ? AND environment_id = ? AND state = ?", args.Release.ProjectID, args.Release.EnvironmentID, transistor.GetState("complete")).Find(&projectExtensions)

		if len(projectExtensions) == 0 {
			log.InfoWithFields("project has no extensions", log.Fields{
				"project_id":     args.Release.ProjectID,
				"environment_id": args.Release.EnvironmentID,
			})
			return nil, fmt.Errorf("no project extensions found")
		}

		projectExtensionsMarshaled, err := json.Marshal(projectExtensions)
		if err != nil {
			return &ReleaseResolver{}, err
		}

		projectExtensionsJsonb = postgres.Jsonb{projectExtensionsMarshaled}
	} else {
		log.Info(fmt.Sprintf("Existing Release. Rolling back %d", args.Release.ID))
		// Rollback
		isRollback = true
		existingRelease := model.Release{}

		if r.DB.Where("id = ?", string(*args.Release.ID)).First(&existingRelease).RecordNotFound() {
			log.ErrorWithFields("Could not find existing release", log.Fields{
				"id": *args.Release.ID,
			})
			return &ReleaseResolver{}, errors.New("Release not found")
		}

		secretsJsonb = existingRelease.Secrets
		servicesJsonb = existingRelease.Services
		projectExtensionsJsonb = existingRelease.ProjectExtensions

		// unmarshal projectExtensionsJsonb and servicesJsonb into project extensions
		err := json.Unmarshal(projectExtensionsJsonb.RawMessage, &projectExtensions)
		if err != nil {
			return &ReleaseResolver{}, errors.New("Could not unmarshal project extensions")
		}

		err = json.Unmarshal(servicesJsonb.RawMessage, &services)
		if err != nil {
			return &ReleaseResolver{}, errors.New("Could not unmarshal services")
		}

		err = json.Unmarshal(secretsJsonb.RawMessage, &secrets)
		if err != nil {
			return &ReleaseResolver{}, errors.New("Could not unmarshal secrets")
		}
	}

	// check if there's a previous release in waiting state that
	// has the same secrets and services signatures
	secretsSha1 := sha1.New()
	secretsSha1.Write(secretsJsonb.RawMessage)
	secretsSig := secretsSha1.Sum(nil)

	servicesSha1 := sha1.New()
	servicesSha1.Write(servicesJsonb.RawMessage)
	servicesSig := servicesSha1.Sum(nil)

	currentReleaseHeadFeature := model.Feature{}

	r.DB.Where("id = ?", args.Release.HeadFeatureID).First(&currentReleaseHeadFeature)

	waitingRelease := model.Release{}

	r.DB.Where("state in (?) and project_id = ? and environment_id = ?", []string{string(transistor.GetState("waiting")),
		string(transistor.GetState("running"))}, args.Release.ProjectID, args.Release.EnvironmentID).Order("created_at desc").First(&waitingRelease)

	wrSecretsSha1 := sha1.New()
	wrSecretsSha1.Write(waitingRelease.Services.RawMessage)
	waitingReleaseSecretsSig := wrSecretsSha1.Sum(nil)

	wrServicesSha1 := sha1.New()
	wrServicesSha1.Write(waitingRelease.Services.RawMessage)
	waitingReleaseServicesSig := wrServicesSha1.Sum(nil)

	waitingReleaseHeadFeature := model.Feature{}

	r.DB.Where("id = ?", waitingRelease.HeadFeatureID).First(&waitingReleaseHeadFeature)

	if bytes.Equal(secretsSig, waitingReleaseSecretsSig) &&
		bytes.Equal(servicesSig, waitingReleaseServicesSig) &&
		strings.Compare(currentReleaseHeadFeature.Hash, waitingReleaseHeadFeature.Hash) == 0 {

		// same release so return
		log.InfoWithFields("Found a waiting release with the same services signature, secrets signature and head feature hash. Aborting", log.Fields{
			"services_sig":      servicesSig,
			"secrets_sig":       secretsSig,
			"head_feature_hash": waitingReleaseHeadFeature.Hash,
		})
		return &ReleaseResolver{}, fmt.Errorf("Found a waiting release with the same properties. Aborting.")
	}

	projectID, err := uuid.FromString(args.Release.ProjectID)
	if err != nil {
		log.InfoWithFields("Couldn't parse projectID", log.Fields{
			"args": args,
		})
		return &ReleaseResolver{}, fmt.Errorf("Couldn't parse projectID")
	}

	headFeatureID, err := uuid.FromString(args.Release.HeadFeatureID)
	if err != nil {
		log.InfoWithFields("Couldn't parse headFeatureID", log.Fields{
			"args": args,
		})
		return &ReleaseResolver{}, fmt.Errorf("Couldn't parse headFeatureID")
	}

	environmentID, err := uuid.FromString(args.Release.EnvironmentID)
	if err != nil {
		log.InfoWithFields("Couldn't parse environmentID", log.Fields{
			"args": args,
		})
		return &ReleaseResolver{}, fmt.Errorf("Couldn't parse environmentID")
	}

	// the tail feature id is the current release's head feature id
	currentRelease := model.Release{}
	tailFeatureID := headFeatureID
	if err = r.DB.Where("state = ? and project_id = ? and environment_id = ?", transistor.GetState("complete"), projectID, environmentID).Find(&currentRelease).Order("created_at desc").Limit(1).Error; err == nil {
		tailFeatureID = currentRelease.HeadFeatureID
	}

	if r.DB.Where("id = ?", projectID).First(&project).RecordNotFound() {
		log.InfoWithFields("project not found", log.Fields{
			"id": projectID,
		})
		return &ReleaseResolver{}, errors.New("Project not found")
	}

	// get all branches relevant for the project
	var branch string
	var projectSettings model.ProjectSettings

	if r.DB.Where("environment_id = ? and project_id = ?", environmentID, projectID).First(&projectSettings).RecordNotFound() {
		log.InfoWithFields("no env project branch found", log.Fields{})
	} else {
		branch = projectSettings.GitBranch
	}

	var environment model.Environment
	if r.DB.Where("id = ?", environmentID).Find(&environment).RecordNotFound() {
		log.InfoWithFields("no env found", log.Fields{
			"id": environmentID,
		})
		return &ReleaseResolver{}, errors.New("Environment not found")
	}

	var headFeature model.Feature
	if r.DB.Where("id = ?", headFeatureID).First(&headFeature).RecordNotFound() {
		log.InfoWithFields("head feature not found", log.Fields{
			"id": headFeatureID,
		})
		return &ReleaseResolver{}, errors.New("head feature not found")
	}

	var tailFeature model.Feature
	if r.DB.Where("id = ?", tailFeatureID).First(&tailFeature).RecordNotFound() {
		log.InfoWithFields("tail feature not found", log.Fields{
			"id": tailFeatureID,
		})
		return &ReleaseResolver{}, errors.New("Tail feature not found")
	}

	var pluginServices []plugins.Service
	pluginServices, err = r.setupServices(services)
	if err != nil {
		return &ReleaseResolver{}, err
	}

	var pluginSecrets []plugins.Secret
	for _, secret := range secrets {
		pluginSecrets = append(pluginSecrets, plugins.Secret{
			Key:   secret.Key,
			Value: secret.Value.Value,
			Type:  secret.Type,
		})
	}

	// Create/Emit Release ProjectExtensions
	willCreateReleaseExtension := false
	for _, projectExtension := range projectExtensions {
		extension := model.Extension{}
		if r.DB.Where("id= ?", projectExtension.ExtensionID).Find(&extension).RecordNotFound() {
			log.ErrorWithFields("extension spec not found", log.Fields{
				"id": projectExtension.ExtensionID,
			})
			return &ReleaseResolver{}, fmt.Errorf("extension spec not found")
		}

		if plugins.Type(extension.Type) == plugins.GetType("workflow") || plugins.Type(extension.Type) == plugins.GetType("deployment") {
			willCreateReleaseExtension = true
			break
		}
	}

	var release model.Release
	if willCreateReleaseExtension == true {
		// Create Release
		release = model.Release{
			State:             transistor.GetState("waiting"),
			StateMessage:      "Release created",
			ProjectID:         projectID,
			EnvironmentID:     environmentID,
			UserID:            uuid.FromStringOrNil(userID),
			HeadFeatureID:     headFeatureID,
			TailFeatureID:     tailFeatureID,
			Secrets:           secretsJsonb,
			Services:          servicesJsonb,
			ProjectExtensions: projectExtensionsJsonb,
			ForceRebuild:      args.Release.ForceRebuild,
			IsRollback:        isRollback,
		}

		r.DB.Create(&release)
	} else {
		return nil, fmt.Errorf("No release extensions found")
	}

	// insert CodeAmp envs
	slugSecret := plugins.Secret{
		Key:   "CODEAMP_SLUG",
		Value: project.Slug,
		Type:  plugins.GetType("env"),
	}
	pluginSecrets = append(pluginSecrets, slugSecret)

	hashSecret := plugins.Secret{
		Key:   "CODEAMP_HASH",
		Value: headFeature.Hash[0:7],
		Type:  plugins.GetType("env"),
	}
	pluginSecrets = append(pluginSecrets, hashSecret)

	timeSecret := plugins.Secret{
		Key:   "CODEAMP_CREATED_AT",
		Value: time.Now().Format(time.RFC3339),
		Type:  plugins.GetType("env"),
	}
	pluginSecrets = append(pluginSecrets, timeSecret)

	// insert Codeflow envs - remove later
	_slugSecret := plugins.Secret{
		Key:   "CODEFLOW_SLUG",
		Value: project.Slug,
		Type:  plugins.GetType("env"),
	}
	pluginSecrets = append(pluginSecrets, _slugSecret)

	_hashSecret := plugins.Secret{
		Key:   "CODEFLOW_HASH",
		Value: headFeature.Hash[0:7],
		Type:  plugins.GetType("env"),
	}
	pluginSecrets = append(pluginSecrets, _hashSecret)

	_timeSecret := plugins.Secret{
		Key:   "CODEFLOW_CREATED_AT",
		Value: time.Now().Format(time.RFC3339),
		Type:  plugins.GetType("env"),
	}
	pluginSecrets = append(pluginSecrets, _timeSecret)

	releaseEvent := plugins.Release{
		IsRollback:  isRollback,
		ID:          release.Model.ID.String(),
		Environment: environment.Key,
		HeadFeature: plugins.Feature{
			ID:         headFeature.Model.ID.String(),
			Hash:       headFeature.Hash,
			ParentHash: headFeature.ParentHash,
			User:       headFeature.User,
			Message:    headFeature.Message,
			Created:    headFeature.Created,
		},
		TailFeature: plugins.Feature{
			ID:         tailFeature.Model.ID.String(),
			Hash:       tailFeature.Hash,
			ParentHash: tailFeature.ParentHash,
			User:       tailFeature.User,
			Message:    tailFeature.Message,
			Created:    tailFeature.Created,
		},
		User: release.User.Email,
		Project: plugins.Project{
			ID:         project.Model.ID.String(),
			Slug:       project.Slug,
			Repository: project.Repository,
		},
		Git: plugins.Git{
			Url:           project.GitUrl,
			Branch:        branch,
			RsaPrivateKey: project.RsaPrivateKey,
		},
		Secrets:  pluginSecrets,
		Services: pluginServices, // ADB Added this
	}

	for _, projectExtension := range projectExtensions {
		extension := model.Extension{}
		if r.DB.Where("id= ?", projectExtension.ExtensionID).Find(&extension).RecordNotFound() {
			log.ErrorWithFields("extension spec not found", log.Fields{
				"id": projectExtension.ExtensionID,
			})
			return &ReleaseResolver{}, errors.New("extension spec not found")
		}

		if plugins.Type(extension.Type) == plugins.GetType("workflow") || plugins.Type(extension.Type) == plugins.GetType("deployment") {
			var headFeature model.Feature
			if r.DB.Where("id = ?", release.HeadFeatureID).First(&headFeature).RecordNotFound() {
				log.ErrorWithFields("head feature not found", log.Fields{
					"id": release.HeadFeatureID,
				})
				return &ReleaseResolver{}, errors.New("head feature not found")
			}

			// create ReleaseExtension
			releaseExtension := model.ReleaseExtension{
				State:              transistor.GetState("waiting"),
				StateMessage:       "",
				ReleaseID:          release.Model.ID,
				FeatureHash:        headFeature.Hash,
				ServicesSignature:  fmt.Sprintf("%x", servicesSig),
				SecretsSignature:   fmt.Sprintf("%x", secretsSig),
				ProjectExtensionID: projectExtension.Model.ID,
				Type:               extension.Type,
			}

			r.DB.Create(&releaseExtension)
		}
	}

	if waitingRelease.State != "" {
		log.Info(fmt.Sprintf("Release is already running, queueing %s", release.Model.ID.String()))
		return &ReleaseResolver{}, fmt.Errorf("Release is already running, queuing %s", release.Model.ID.String())
	} else {
		release.State = transistor.GetState("running")
		release.StateMessage = "Running Release"
		release.Started = time.Now()
		r.DB.Save(&release)

		r.Events <- transistor.NewEvent(transistor.EventName("release"), transistor.GetAction("create"), releaseEvent)
		return &ReleaseResolver{DBReleaseResolver: &db_resolver.ReleaseResolver{DB: r.DB, Release: release}}, nil
	}
}

func (r *ReleaseResolverMutation) StopRelease(ctx context.Context, args *struct{ ID graphql.ID }) (*ReleaseResolver, error) {
	userID, err := auth.CheckAuth(ctx, []string{})
	if err != nil {
		return &ReleaseResolver{}, err
	}

	var user model.User

	r.DB.Where("id = ?", userID).Find(&user)

	var release model.Release
	var releaseExtensions []model.ReleaseExtension

	r.DB.Where("release_id = ?", args.ID).Find(&releaseExtensions)
	if len(releaseExtensions) < 1 {
		log.WarnWithFields("No release extensions found for release: %s", log.Fields{
			"id": args.ID,
		})
	}

	if r.DB.Where("id = ?", args.ID).Find(&release).RecordNotFound() {
		log.WarnWithFields("Release not found", log.Fields{
			"id": args.ID,
		})

		return nil, errors.New("Release Not Found")
	}

	release.State = transistor.GetState("canceled")
	release.StateMessage = fmt.Sprintf("Release canceled by %s", user.Email)
	r.DB.Save(&release)

	for _, releaseExtension := range releaseExtensions {
		var projectExtension model.ProjectExtension
		if r.DB.Where("id = ?", releaseExtension.ProjectExtensionID).Find(&projectExtension).RecordNotFound() {
			log.WarnWithFields("Associated project extension not found", log.Fields{
				"id":                   args.ID,
				"release_extension_id": releaseExtension.ID,
				"project_extension_id": releaseExtension.ProjectExtensionID,
			})

			return nil, errors.New("Project Extension Not Found")
		}

		// find associated ProjectExtension Extension
		var extension model.Extension
		if r.DB.Where("id = ?", projectExtension.ExtensionID).Find(&extension).RecordNotFound() {
			log.WarnWithFields("Associated extension not found", log.Fields{
				"id":                   args.ID,
				"release_extension_id": releaseExtension.ID,
				"project_extension_id": releaseExtension.ProjectExtensionID,
				"extension_id":         projectExtension.ExtensionID,
			})

			return nil, errors.New("Extension Not Found")
		}

		if releaseExtension.State == transistor.GetState("waiting") {
			releaseExtensionEvent := plugins.ReleaseExtension{
				ID:      releaseExtension.ID.String(),
				Project: plugins.Project{},
				Release: plugins.Release{
					ID: releaseExtension.ReleaseID.String(),
				},
				Environment: "",
			}

			// Update the release extension
			event := transistor.NewEvent(transistor.EventName(fmt.Sprintf("release:%s", extension.Key)), transistor.GetAction("create"), releaseExtensionEvent)
			event.State = transistor.GetState("canceled")
			event.StateMessage = fmt.Sprintf("Deployment Stopped By User %s", user.Email)
			r.Events <- event
		}
	}

	return &ReleaseResolver{DBReleaseResolver: &db_resolver.ReleaseResolver{DB: r.DB, Release: release}}, nil
}

func (r *ReleaseResolverMutation) setupServices(services []model.Service) ([]plugins.Service, error) {
	var pluginServices []plugins.Service
	for _, service := range services {
		var spec model.ServiceSpec
		if r.DB.Where("service_id = ?", service.Model.ID).First(&spec).RecordNotFound() {
			log.WarnWithFields("servicespec not found", log.Fields{
				"service_id": service.Model.ID,
			})
			return []plugins.Service{}, fmt.Errorf("ServiceSpec not found")
		}

		pluginServices = AppendPluginService(pluginServices, service, spec)
	}

	return pluginServices, nil
}
