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
	Environment       string `envconfig:"ENV" default:"local"`
	S3BucketName      string `envconfig:"S3_BUCKET_NAME"`
	S3BucketUploadKey string `envconfig:"S3_BUCKET_UPLOAD_KEY"`
	S3BucketRegion    string `envconfig:"S3_BUCKET_REGION"`
	RedisUrl          string `envconfig:"REDIS_URL" default:"redis.svc.cluster.local:6379"`
	QueueUrl          string `envconfig:"QUEUE_URL"`
}

func Load() (Config, error) {
	err := LoadEnvironment()

	cfg := Config{}
	err = envconfig.Process("settings", &cfg)
	return cfg, err
}
