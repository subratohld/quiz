package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env           string                    `yaml:"env"`
	Service       string                    `yaml:"service"`
	Logging       Logging                   `yaml:"logging"`
	Databases     Databases                 `yaml:"db"`
	NetworkConfig map[string]*NetworkConfig `yaml:"services"`
}

type Logging struct {
	Level string `yaml:"level"`
}

type Databases struct {
	Postgres Database `yaml:"postgres"`
}

type Database struct {
	Host   string `yaml:"host"`
	DBName string `yaml:"dbName"`
	Secret Secret `yaml:"secret"`
}

type Secret struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type NetworkConfig struct {
	ServiceName    string `yaml:"serviceName"`
	Port           int    `yaml:"port"`
	DeploymentType string `yaml:"deploymentType"`
}

var (
	_config    *Config
	configOnce sync.Once
)

func InitAndGetConfig() *Config {
	configOnce.Do(func() {
		initConfig()
	})

	return _config
}

func GetConfig() *Config {
	return _config
}

func initConfig() {
	var cfg Config
	if err := readFileAndUnmarshal(serviceConfigFile, &cfg); err != nil {
		return
	}

	var postgresSecret Database
	if err := readFileAndUnmarshal(postgresSecretFile, &postgresSecret); err != nil {
		return
	}

	var networkConfig Config
	if err := readFileAndUnmarshal(networkConfigFile, &networkConfig); err != nil {
		return
	}

	cfg.Databases.Postgres.Secret = postgresSecret.Secret
	cfg.NetworkConfig = networkConfig.NetworkConfig

	_config = &cfg
}

func readFileAndUnmarshal(filePath string, target any) error {
	byteContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err := yaml.Unmarshal(byteContent, target); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (c *Config) GetPort(serviceName string) int {
	networkConfig := c.NetworkConfig[serviceName]
	if networkConfig != nil {
		return networkConfig.Port
	}

	return 0
}
