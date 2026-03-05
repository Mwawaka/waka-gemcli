package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ConfigLoader interface {
	Load() (*Config, error)
}

type Config struct {
	APIKey string
}

type ViperLoader struct {
	FileType string
}

type GoLoader struct {
}

func (v *ViperLoader) Load() (*Config, error) {
	viper.SetConfigFile(v.FileType)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	apiKey := viper.GetString("API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("missing API key")
	}
	return &Config{
		APIKey: apiKey,
	}, nil
}

func (g *GoLoader) Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("missing API key")
	}
	return &Config{
		APIKey: apiKey,
	}, nil
}

// func GetConfig(cfgLoader ConfigLoader) (*Config, error) {
// 	return cfgLoader.Load()

// }
