package config

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const DEFAULT_CONFIG_PATH = "$HOME/.jarvis/config.yaml"

type (
	Configuration struct {
		InventoryPath string `mapstructure:"inventory_path"`
	}

	ConfigurationReader struct {
		path string
	}
)

func InitConfigurationReader(path string) *ConfigurationReader {
	return &ConfigurationReader{path}
}

func (c *ConfigurationReader) ParseConfig() (*Configuration, error) {
	dir, file := filepath.Split(DEFAULT_CONFIG_PATH)
	//viper don't care about the file type
	//we must remove it from the file name
	file = fileNameWithoutExtension(file)
	viper.SetConfigType("yaml")
	viper.SetConfigName(file)
	viper.AddConfigPath(dir)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	conf := &Configuration{}
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (c *ConfigurationReader) GetConfigPath() string {
	return c.path
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
