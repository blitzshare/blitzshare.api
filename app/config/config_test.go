package config_test

import (
	"os"
	"testing"

	"blitzshare.fileshare.api/app/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	setUp()
	cfg, err := config.Load()

	assert.Nil(t, err, "Unable to log the config")
	assert.Equal(t, "local", cfg.Settings.Environment)
	assert.Equal(t, 8000, cfg.Server.Port)

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
