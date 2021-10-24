package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server   Server
	Settings Settings
}

type Server struct {
	Port int `envconfig:"PORT"`
}

type Settings struct {
	Environment string `envconfig:"ENV" default:"local"`
}

func Load() (Config, error) {
	err := LoadEnvironment()

	cfg := Config{}
	err = envconfig.Process("settings", &cfg)
	return cfg, err
}

func LoadEnvironment() {
	panic("unimplemented")
}
