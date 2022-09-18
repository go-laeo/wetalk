package config

import "github.com/caarlos0/env/v6"

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
	SecretToken string `env:"SECRET_TOKEN"`
	SiteName    string `env:"SITE_NAME"`
}

// FromEnv parses configuration from environment variables.
func FromEnv() *Config {
	c := &Config{}
	err := env.Parse(c)
	if err != nil {
		panic(err)
	}
	return c
}
