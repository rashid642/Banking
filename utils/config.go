package utils

import (
	"github.com/spf13/viper"
)

// The values are read from viper from a config file or anvironment varibles
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// Load config read configurations from file or env variables if provided
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
