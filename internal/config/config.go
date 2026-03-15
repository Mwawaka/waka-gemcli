package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var ErrConfigNotFound = fmt.Errorf("config file not found")

type Config struct {
	APIKey string
	Model  string
}

type ViperLoader struct {
	configPath string
}

func NewViperLoader() (*ViperLoader, error) {
	var viperLoader ViperLoader
	home, err := os.UserHomeDir()

	if err != nil {
		return nil, fmt.Errorf("couldn't find home directory: %w", err)
	}

	viperLoader.configPath = filepath.Join(home, ".config", "gemcli", "config.yaml")
	viperLoader.configure()

	return &viperLoader, nil

}

func (v *ViperLoader) configure() {
	viper.SetConfigFile(v.configPath)
}

func (v *ViperLoader) Load() (*Config, error) {
	var pathErr *fs.PathError

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &pathErr) {
			return nil, ErrConfigNotFound
		}
		return nil, err
	}

	apiKey := viper.GetString("GOOGLE_API_KEY")
	model := viper.GetString("MODEL")

	if apiKey == "" || model == "" {
		return nil, fmt.Errorf("missing configuration variables")
	}
	return &Config{
		APIKey: apiKey,
		Model:  model,
	}, nil
}

func (v *ViperLoader) Set(key, value string) error {
	dir := filepath.Dir(v.configPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}

	viper.Set(key, value)

	if err := viper.WriteConfig(); err != nil {
		if err := viper.SafeWriteConfig(); err != nil {
			return err
		}
	}

	return nil
}

func (v *ViperLoader) Get(key string) (string, error) {
	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("reading configuration: %w", err)
	}

	return viper.GetString(key), nil
}
