package test

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/codeamp/circuit/plugins"
	"github.com/codeamp/circuit/plugins/codeamp/model"
	log "github.com/codeamp/logger"
	"github.com/codeamp/transistor"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const ValidUUID = "123e4567-e89b-12d3-a456-426655440000"
const InvalidUUID = "123e4567-z89b-12d3-a456-426655440000"

// Setup logic common to all Resolver tests
// Load viper config, parse plugins, and setup log level/format
// Configs in this case are loaded from a test yml file
func SetupResolverTest(migrators []interface{}) (*gorm.DB, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("../../../configs/circuit.test.yml")

	setupViperEnvs()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	db, err := setupPostgresDB()
	if err != nil {
		return nil, err
	}

	for _, migrator := range migrators {
		db.AutoMigrate(migrator)
	}

	configLogLevel()
	configLogFormat()

	events := []interface{}{
		plugins.ProjectExtension{},
		plugins.ReleaseExtension{},
		plugins.Release{},
		plugins.ReleaseExtension{},
	}
	for _, i := range events {
		transistor.EventRegistry[reflect.TypeOf(i).String()] = i
	}

	return db, nil
}

// Generates a fake JWT token for testing purposes
// Use for testing graphql resolvers
func ResolverAuthContext() context.Context {
	return BuildAuthContext("11075553-5309-494B-9085-2D79A6ED1EB3", "foo@example.com", []string{"admin"})
}

// Generates a fake JWT token for testing purposes
// Use for testing graphql resolvers
// Allows more flexibility in creating a prepopulated context
func BuildAuthContext(userID string, email string, permissionsList []string) context.Context {
	authContext := context.WithValue(context.Background(), "jwt", model.Claims{
		UserID:      userID,
		Email:       email,
		Permissions: permissionsList,
	})

	return authContext
}

func setupPostgresDB() (*gorm.DB, error) {
	log.DebugWithFields("Setup Postgres DB Connection for Tests",
		log.Fields{
			"host":     viper.GetString("plugins.codeamp.postgres.host"),
			"port":     viper.GetString("plugins.codeamp.postgres.port"),
			"user":     viper.GetString("plugins.codeamp.postgres.user"),
			"dbname":   viper.GetString("plugins.codeamp.postgres.dbname"),
			"sslmode":  viper.GetString("plugins.codeamp.postgres.sslmode"),
			"password": "*",
		},
	)

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		viper.GetString("plugins.codeamp.postgres.host"),
		viper.GetString("plugins.codeamp.postgres.port"),
		viper.GetString("plugins.codeamp.postgres.user"),
		viper.GetString("plugins.codeamp.postgres.dbname"),
		viper.GetString("plugins.codeamp.postgres.sslmode"),
		viper.GetString("plugins.codeamp.postgres.password"),
	))

	// db.LogMode(true)
	return db, err
}

// Setup logic common to all Plugin tests
// Load viper config, parse plugins, and setup log level/format and instantiate transistor
// Configs in this case are loaded inline from the plugin test itself
func SetupPluginTest(viperConfig []byte) (*transistor.Transistor, error) {
	setupViperConfig(viperConfig)

	pluginConfig := viper.GetStringMap("plugins")
	enabledPluginNames := make([]string, 0, len(pluginConfig))

	for pluginName, _ := range pluginConfig {
		enabledPluginNames = append(enabledPluginNames, pluginName)
	}

	config := transistor.Config{
		Plugins:        pluginConfig,
		EnabledPlugins: enabledPluginNames,
	}

	configLogLevel()
	configLogFormat()

	return transistor.NewTestTransistor(config)
}

func setupViperConfig(viperConfig []byte) {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(viperConfig))

	setupViperEnvs()
}

func setupViperEnvs() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("CODEAMP")
	viper.AutomaticEnv()
}

func configLogLevel() {
	if _logLevel := viper.GetString("log_level"); _logLevel != "" {
		logLevel, err := log.ParseLevel(_logLevel)

		if err != nil {
			log.Fatal(err)
		}

		log.SetLogLevel(logLevel)
	}
}

func configLogFormat() {
	if logFormat := viper.GetString("log_format"); logFormat != "" {
		switch strings.ToLower(logFormat) {
		case "standard":
			break
		case "json":
			fallthrough
		default:
			log.SetLogFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
		}
	}
}
