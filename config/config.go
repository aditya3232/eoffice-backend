package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var ENV *config

type config struct {
	JWT_KEY            string `mapstructure:"JWT_KEY"`
	DB1_USERNAME       string `mapstructure:"DB1_USERNAME"`
	DB1_PASSWORD       string `mapstructure:"DB1_PASSWORD"`
	DB1_HOST           string `mapstructure:"DB1_HOST"`
	DB1_PORT           string `mapstructure:"DB1_PORT"`
	DB1_DATABASE       string `mapstructure:"DB1_DATABASE"`
	REDIS_USERNAME     string `mapstructure:"REDIS_USERNAME"`
	REDIS_PASSWORD     string `mapstructure:"REDIS_PASSWORD"`
	REDIS_HOST         string `mapstructure:"REDIS_HOST"`
	REDIS_PORT         string `mapstructure:"REDIS_PORT"`
	REDIS_DATABASE     string `mapstructure:"REDIS_DATABASE"`
	ELASTICSEARCH_HOST string `mapstructure:"ELASTICSEARCH_HOST"`
	ELASTICSEARCH_PORT string `mapstructure:"ELASTICSEARCH_PORT"`
	DEBUG              int    `mapstructure:"DEBUG"`
}

func init() {
	ENV = LoadConfig()
}

func LoadConfig() *config {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get current working directory: %w", err))
	}

	// Move up directories until we find one that contains a file named "go.mod"
	for {
		if _, err := os.Stat(cwd + "/go.mod"); err == nil {
			break
		}
		cwd = filepath.Dir(cwd)
	}

	viper.AddConfigPath(cwd + "/config")
	viper.SetConfigType("env")

	// if testing is running, use .env.test
	if flag.Lookup("test.v") != nil {
		viper.SetConfigName(".env.test")
	} else {
		viper.SetConfigName(".env")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read env file: %w", err))
	}

	config := &config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("failed to unmarshal env variables: %w", err))
	}

	return config
}
