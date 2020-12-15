package config

import (
	"io/ioutil"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config contains the application configuration
type Config struct {
	HTTP struct {
		Port int `yaml:"port"`
	}
	DB struct {
		DBtype          string `yaml:"dbtype"`
		DBurl           string `yaml:"dburl"`
		MaxIdleConns    int    `yaml:"maxIdleConns"`
		MaxOpenConns    int    `yaml:"maxOpenConns"`
		ConnMaxLifetime string `yaml:"connMaxLifetime"`
	}
	SECURITY struct {
		Resource  string `yaml:"resource"`
		SecretKey string `yaml:"secretKey"`
	}
}

var config *Config
var configPath string

// Get return the configurations
func Get() *Config {
	if config != nil {
		return config
	}

	configPath = getConfigFilePath()
	yml, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("Attention! Can't read configuration file!")
	}

	config = &Config{}
	err = yaml.UnmarshalStrict(yml, config)
	if err != nil {
		panic("Attention! Invalid config format!")
	}

	return config
}

// this function can return config files for diff env in the future
func getConfigFilePath() string {
	_, file, _, ok := runtime.Caller(0)

	if !ok {
		panic("Attention! Can't get configuration file path!")
	}

	idx := strings.LastIndexByte(file, '/')
	return file[:idx] + "/config.yaml"
}
