package util

import "github.com/spf13/viper"

type Config struct {
	DB_DRIVER      string `mapstructure:"DB_DRIVER"`
	DB_SOURCE      string `mapstructure:"DB_SOURCE"`
	SERVER_ADDRESS string `mapstructure:"SERVER_ADDRESS"`

	USER_AUTH_ACCESS_JWT_SECRETE_KEY  string `mapstructure:"USER_AUTH_ACCESS_JWT_SECRETE_KEY"`
	USER_AUTH_REFRESH_JWT_SECRETE_KEY string `mapstructure:"USER_AUTH_REFRESH_JWT_SECRETE_KEY"`
	USER_AUTH_JWT_EXPIRATION          string `mapstructure:"USER_AUTH_JWT_EXPIRATION"`
	USER_AUTH_CUSTOM_KEY_SECRETE      string `mapstructure:"USER_AUTH_CUSTOM_KEY_SECRETE"`
}

// Loadconfig reads from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
