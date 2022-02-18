package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server   Server
	Settings Settings
	ClientId string
}

type Server struct {
	Port int `envconfig:"PORT"`
}

type Settings struct {
	Environment          string `envconfig:"ENV" default:"local"`
	RedisUrl             string `envconfig:"REDIS_URL" default:"redis.svc.cluster.local:6379"`
	QueueUrl             string `envconfig:"QUEUE_URL"`
	KeyStoreDbConnection string `envconfig:"KEYSTORE_DB_CONNECTION"`
}

func Load() (Config, error) {
	err := LoadEnvironment()

	cfg := Config{
		ClientId: "blitzshare.api",
	}
	err = envconfig.Process("settings", &cfg)
	return cfg, err
}
