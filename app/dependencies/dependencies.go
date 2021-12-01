package dependencies

import (
	"blitzshare.api/app/config"
	"blitzshare.api/app/server/services/registry"
)

type Dependencies struct {
	Config   config.Config
	Registry registry.Registry
}

func NewDependencies(config config.Config) (*Dependencies, error) {
	return &Dependencies{
		Config:   config,
		Registry: registry.NewRegistry(config.Settings.RedisUrl),
	}, nil
}
