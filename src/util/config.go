package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	DBEngine             string        `mapstructure:"DB_ENGINE"`
	HTTPAddress          string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	// Setup path and filename with extension of configuration file
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Auto check and load configuration file to viper if exist
	viper.AutomaticEnv()

	// Read environment variable
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Load environment variable to config
	err = viper.Unmarshal(&config)
	return
}
