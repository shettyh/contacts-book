package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

const (
	envPrefix = "CB"
)

var configInstance struct {
	once     sync.Once
	instance *Config
}

type Config struct {
	DbHost     string `required: "true"` // CB_DBHOST
	DbPort     int    `default:3306`
	DbUser     string `default:"root"`
	DbPassword string
	DbName     string `default:"contactsbook"`
}

func GetInstance() *Config {
	configInstance.once.Do(func() {
		configInstance.instance = new(Config)
		err := envconfig.Process(envPrefix, &configInstance.instance)
		if err != nil {
			panic(err)
		}
	})
	return configInstance.instance
}
