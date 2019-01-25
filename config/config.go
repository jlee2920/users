package config

import (
	"github.com/caarlos0/env"
)

// Conf holds the environment configurations taken from docker-compose.yml
type Conf struct {
	DBUser     	string 	`env:"DB_USER"`
	DBPassword	string	`env:"DB_PASSWORD"`
	DBHost		string	`env:"DB_HOST"`
	DBPort		string	`env:"DB_PORT"`
	DBName		string	`env:"DB_NAME"`
	DBType		string 	`env:"DB_TYPE"`
}

// InitEnv grabs the environment variables found within the docker-compose.yml file
func InitializeConfig() Conf {
	var conf Conf
	// Config is a global configuration that is used within qoqbot
	if err := env.Parse(&conf); err != nil {
		panic(err)
	}
	return conf
}