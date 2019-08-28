package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

const (
	// All the util env's should start with this prefix
	envPrefix = "CB"
)

// configInstance singleton holder for util.
var configInstance struct {
	once     sync.Once
	instance *Config
}

type Config struct {
	// DB configs
	DbType     string `default:"mysql"` // CB_DBTYPE
	DbHost     string `required:"true"` // CB_DBHOST
	DbPort     int    `default:"3306"`  // CB_DBPORT
	DbUser     string `default:"root"`  // CB_DBUSER
	DbPassword string // CB_DBPASSWORD
	DbName     string `default:"contactsbook"` // CB_DBNAME

	// Server util
	Port int `default:"80"`
}

// GetInstance will create or get the util singleton instance.
func GetInstance() *Config {
	configInstance.once.Do(func() {
		configInstance.instance = new(Config)
		err := envconfig.Process(envPrefix, configInstance.instance)
		// exit if loading util is failed.
		if err != nil {
			log.Fatalf("failed to load the configs, %v", err)
		}
	})
	return configInstance.instance
}
