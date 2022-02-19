package dependencies

import (
	"blitzshare.api/app/config"
	"blitzshare.api/app/services/key"

	"blitzshare.api/app/services/event"
	"blitzshare.api/app/services/random"
	"blitzshare.api/app/services/registry"
)

type Dependencies struct {
	Config      config.Config
	Registry    registry.Registry
	EventEmit   event.EventEmit
	ApiKeychain key.ApiKeychain
	Rnd         random.Rnd
}

func NewDependencies(config config.Config) (*Dependencies, error) {
	return &Dependencies{
		Config:      config,
		Registry:    registry.New(config.Settings.RedisUrl),
		EventEmit:   event.New(),
		ApiKeychain: key.New(config),
		Rnd:         random.NewRnd(),
	}, nil
}
