package db_resolver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/codeamp/circuit/plugins/codeamp/auth"
	"github.com/codeamp/circuit/plugins/codeamp/model"
	log "github.com/codeamp/logger"
	"github.com/codeamp/transistor"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// ReleaseResolver resolver for Release
type ReleaseResolver struct {
	model.Release
	DB *gorm.DB
}

// ID
func (r *ReleaseResolver) ID() graphql.ID {
	return graphql.ID(r.Release.Model.ID.String())
}

// Project
func (r *ReleaseResolver) Project() *ProjectResolver {
	var project model.Project

	r.DB.Model(r.Release).Related(&project)

	return &ProjectResolver{DB: r.DB, Project: project}
}

// User
func (r *ReleaseResolver) User() *UserResolver {
	var user model.User

	r.DB.Model(r.Release).Related(&user)

	return &UserResolver{DB: r.DB, User: user}
}

// Artifacts
func (r *ReleaseResolver) Artifacts(ctx context.Context) (model.JSON, error) {
	artifacts := []transistor.Artifact{}
	var releaseExtensions []model.ReleaseExtension

	if _, err := auth.CheckAuth(ctx, []string{}); err != nil {
		return model.JSON{[]byte("[]")}, err
	}

	isAdmin := false
	if _, err := auth.CheckAuth(ctx, []string{"admin"}); err == nil {
		isAdmin = true
	}

	r.DB.Where("release_id = ?", r.Model.ID).Find(&releaseExtensions)

	for _, releaseExtension := range releaseExtensions {
		var _artifacts []transistor.Artifact

		projectExtension := model.ProjectExtension{}
		if r.DB.Unscoped().Where("id = ?", releaseExtension.ProjectExtensionID).Find(&projectExtension).RecordNotFound() {
			log.InfoWithFields("project extensions not found", log.Fields{
				"id": releaseExtension.ProjectExtensionID,
				"release_extension_id": releaseExtension.Model.ID,
			})
			return model.JSON{[]byte("[]")}, errors.New("release extension not found")
		}

		extension := model.Extension{}
		if r.DB.Where("id= ?", projectExtension.ExtensionID).Find(&extension).RecordNotFound() {
			log.InfoWithFields("extension not found", log.Fields{
				"id": projectExtension.Model.ID,
				"release_extension_id": releaseExtension.Model.ID,
			})
			return model.JSON{[]byte("[]")}, errors.New("project extension not found")
		}

		err := json.Unmarshal(releaseExtension.Artifacts.RawMessage, &_artifacts)
		if err != nil {
			log.InfoWithFields(err.Error(), log.Fields{
				"input": releaseExtension.Artifacts.RawMessage,
			})
		} else {
			for _, artifact := range _artifacts {
				artifact.Source = extension.Key
				artifacts = append(artifacts, artifact)
			}
		}
	}

	for i, artifact := range artifacts {
		if !isAdmin && artifact.Secret {
			artifacts[i].Value = ""
		}
	}

	marshalledArtifacts, err := json.Marshal(artifacts)
	if err != nil {
		log.InfoWithFields(err.Error(), log.Fields{
			"input": artifacts,
		})
		return model.JSON{[]byte("[]")}, err
	}

	return model.JSON{json.RawMessage(marshalledArtifacts)}, nil
}

// HeadFeature
func (r *ReleaseResolver) HeadFeature() *FeatureResolver {
	var feature model.Feature
	r.DB.Where("id = ?", r.Release.HeadFeatureID).First(&feature)
	return &FeatureResolver{DB: r.DB, Feature: feature}
}

// TailFeature
func (r *ReleaseResolver) TailFeature() *FeatureResolver {
	var feature model.Feature

	r.DB.Where("id = ?", r.Release.TailFeatureID).First(&feature)

	return &FeatureResolver{DB: r.DB, Feature: feature}
}

// State
func (r *ReleaseResolver) State() string {
	return string(r.Release.State)
}

// ReleaseExtensions
func (r *ReleaseResolver) ReleaseExtensions() []*ReleaseExtensionResolver {
	var rows []model.ReleaseExtension
	var results []*ReleaseExtensionResolver

	r.DB.Where("release_extensions.release_id = ?", r.Release.Model.ID).Joins(`INNER JOIN project_extensions ON release_extensions.project_extension_id = project_extensions.id 
		INNER JOIN extensions ON project_extensions.extension_id = extensions.id`).Order(`
			CASE extensions.type
				WHEN 'workflow' THEN 1
				WHEN 'deployment' THEN 2
				ELSE 3
			END, extensions.key ASC`).Find(&rows)

	for _, releaseExtension := range rows {
		results = append(results, &ReleaseExtensionResolver{DB: r.DB, ReleaseExtension: releaseExtension})
	}

	return results
}

// StateMessage
func (r *ReleaseResolver) StateMessage() string {
	return r.Release.StateMessage
}

// Environment
func (r *ReleaseResolver) Environment() (*EnvironmentResolver, error) {
	var environment model.Environment
	if r.DB.Where("id = ?", r.Release.EnvironmentID).First(&environment).RecordNotFound() {
		log.InfoWithFields("environment not found", log.Fields{
			"releaseID": r.Release.Model.ID,
		})
		return nil, fmt.Errorf("Environment not found.")
	}
	return &EnvironmentResolver{DB: r.DB, Environment: environment}, nil
}
