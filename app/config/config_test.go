package config_test

import (
	"os"
	"testing"

	"blitzshare.api/app/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	setUp()
	cfg, err := config.Load()

	assert.Nil(t, err, "Unable to log the config")
	assert.Equal(t, "local", cfg.Settings.Environment)
	assert.Equal(t, 8000, cfg.Server.Port)
	assert.Equal(t, "redis.svc.cluster.local", cfg.Settings.RedisUrl)

	tearDown()
}

func setUp() {
	_ = os.Setenv("PORT", "8000")
	_ = os.Setenv("REDIS_URL", "redis.svc.cluster.local")
	_ = os.Setenv("ENV", "local")
}

func tearDown() {
	_ = os.Unsetenv("PORT")
	_ = os.Unsetenv("ENV")
	_ = os.Unsetenv("REDIS_URL")
}
