package dependencies

import (
	"blitzshare.api/app/config"

	"blitzshare.api/app/services/event"
	"blitzshare.api/app/services/random"
	"blitzshare.api/app/services/registry"
)

type Dependencies struct {
	Config    config.Config
	Registry  registry.Registry
	EventEmit event.EventEmit
	Rnd       random.Rnd
}

func NewDependencies(config config.Config) (*Dependencies, error) {
	return &Dependencies{
		Config:    config,
		Registry:  registry.NewRegistry(config.Settings.RedisUrl),
		EventEmit: event.NewEventEmit(),
		Rnd:       random.NewRnd(),
	}, nil
}
