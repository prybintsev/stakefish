package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	IsKubernetes   bool   `mapstructure:"IS_KUBERNETES"`
	Port           int    `mapstructure:"PORT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBName         string `mapstructure:"DB_NAME"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBPort         int    `mapstructure:"DB_PORT"`
	MigrationsPath string `mapstructure:"MIGRATIONS_PATH"`
}

func Init(logEntry *logrus.Entry) AppConfig {
	var config AppConfig
	viper.AutomaticEnv()
	bindEnvironmentVariables(logEntry)
	if err := viper.Unmarshal(&config); err != nil {
		logEntry.WithError(err).Fatal("Unable to decode into struct")
	}

	return config
}

func bindEnvironmentVariables(logEntry *logrus.Entry) {
	environmentVariables := []string{
		"IS_KUBERNETES",
		"PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASSWORD",
		"MIGRATIONS_PATH",
	}

	viper.SetDefault("PORT", 3000)
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_NAME", "stakefish")
	viper.SetDefault("MIGRATIONS_PATH", "internal/db/migrations/scripts")

	for _, environmentVariable := range environmentVariables {
		if err := viper.BindEnv(environmentVariable); err != nil {
			logEntry.WithError(err).WithField("envVariable", environmentVariable).
				Fatal("Unable to bind environment variable")
		}
	}
}
