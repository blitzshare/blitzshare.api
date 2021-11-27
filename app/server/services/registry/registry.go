package registry

import (
	"errors"

	deps "blitzshare.api/app/dependencies"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var client *redis.Client

func getClient(d *deps.Dependencies) *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     d.Config.Settings.RedisUrl,
			Password: "",
		})
		pong, _ := client.Ping().Result()
		log.Infoln("getClient pong", pong)
	}
	return client
}

func get(d *deps.Dependencies, key string) (string, error) {
	client := getClient(d)
	return client.Get(key).Result()
}

func GetPeerMultiAddr(d *deps.Dependencies, pass string) (string, error) {
	result, err := get(d, pass)
	if err == nil {
		return result, err
	}
	return "", errors.New("PeerNotFoundError")
}
