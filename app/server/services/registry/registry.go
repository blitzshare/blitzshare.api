package registry

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var client *redis.Client

func getClient(addr string) *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
		})
		pong, _ := client.Ping().Result()
		log.Infoln("getClient pong", pong)
	}
	return client
}

type Registry interface {
	Get(otp string) (string, error)
}

type RegistryIml struct {
	redisUrl string
}

func NewRegistry(redisUrl string) Registry {

	return &RegistryIml{redisUrl: redisUrl}
}

func (impl *RegistryIml) Get(otp string) (string, error) {
	client := getClient(impl.redisUrl)
	return client.Get(otp).Result()
}
