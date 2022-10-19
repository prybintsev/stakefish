package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	IsKubernetes bool `mapstructure:"IS_KUBERNETES"`
	Port         int  `mapstructure:"PORT"`
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
	}

	viper.SetDefault("PORT", 3000)

	for _, environmentVariable := range environmentVariables {
		if err := viper.BindEnv(environmentVariable); err != nil {
			logEntry.WithError(err).WithField("envVariable", environmentVariable).
				Fatal("Unable to bind environment variable")
		}
	}
}
