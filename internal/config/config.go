package config

type AppConfig struct {
	IsKubernetes bool `mapstructure:"IS_KUBERNETES"`
}
