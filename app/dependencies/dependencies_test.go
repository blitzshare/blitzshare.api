package dependencies_test

import (
	"fmt"
	"os"
	"testing"

	"blitzshare.fileshare.api/app/dependencies"

	"blitzshare.fileshare.api/app/config"
	"github.com/stretchr/testify/assert"
)

func TestDependencies(t *testing.T) {
	setUp()

	cfg, err := config.Load()
	fmt.Printf("%v", cfg.Server.Port)
	assert.Nil(t, err, "config should load")

	deps, err := dependencies.NewDependencies(cfg)
	assert.Nil(t, err, "dependencies should be initialised")
	assert.NotNilf(t, deps.Config, "config should be loaded in dependencies")

	tearDown()
}

func setUp() {
	_ = os.Setenv("PORT", "8000")
	_ = os.Setenv("ENV", "local")
}

func tearDown() {
	_ = os.Unsetenv("PORT")
	_ = os.Unsetenv("ENV")
}
