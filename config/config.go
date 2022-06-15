package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbStr     string `mapstructure:"DB_STRING"`
	Port      uint32 `mapstructure:"PORT"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
	Env       string `mapstructure:"ENV"`
}

var configInstance *Config

func loadConfig(path string) (err error) {
	config := Config{}
	viper.SetDefault("PORT", 3000)
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// tell viper to automatically override values that it has read from config file with the values of the corresponding environment variables if they exist.
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if err == nil {
		configInstance = &config
	}

	return
}

func GetConfig(path string) (*Config, error) {
	if configInstance != nil {
		return configInstance, nil
	}

	err := loadConfig(path)

	return configInstance, err
}
