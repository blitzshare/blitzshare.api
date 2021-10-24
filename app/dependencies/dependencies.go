package dependencies

import "github.com/blitzshare/blitzshare.fileshare.api/app/config"

type Dependencies struct {
	Config config.Config
}

func NewDependencies(config config.Config) (*Dependencies, error) {
	return &Dependencies{Config: config}, nil
}
