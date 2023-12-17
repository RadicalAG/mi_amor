package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig   `mapstructure:"server"`
	DB     DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
	Username string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Uri      string `mapstructure:"uri"`
}

func Init() *Config {
	// Initialize Viper
	viper.SetConfigName("config")   // Name of config file (without extension)
	viper.SetConfigType("yml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // Optionally look for config in the working directory
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &config
}
