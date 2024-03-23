package config

import "github.com/spf13/viper"

type EnvConfig struct {
	ServerPort string `mapstructure:"SERVER_PORT"`

	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`

	KeycloakBaseURL      string `mapstructure:"KEYCLOAK_BASE_URL"`
	KeycloakRealm        string `mapstructure:"KEYCLOAK_REALM"`
	KeycloakClientId     string `mapstructure:"KEYCLOAK_CLIENT_ID"`
	KeycloakUsername     string `mapstructure:"KEYCLOAK_USERNAME"`
	KeycloakClientSecret string `mapstructure:"KEYCLOAK_CLIENT_SECRET"`
}

var Env EnvConfig

func LoadEnv(configPath string) error {
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.UnmarshalExact(&Env); err != nil {
		return err
	}

	return nil
}
