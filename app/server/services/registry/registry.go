package registry

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var client *redis.Client

const (
	P2pPeersDb         = 0
	P2pBootstrapNodeDb = 1
	BootstrapNode      = "BootstrapNode"
)

type Registry interface {
	GetOtp(otp string) (string, error)
	GetNodeConfig() (string, error)
}

type RegistryIml struct {
	redisUrl string
}

func NewRegistry(redisUrl string) Registry {
	return &RegistryIml{redisUrl: redisUrl}
}

func (impl *RegistryIml) getClient(db int) *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     impl.redisUrl,
			Password: "",
			DB:       db,
		})
		pong, _ := client.Ping().Result()
		log.Infoln("getClient pong", pong)
	}
	return client
}

func (impl *RegistryIml) GetOtp(otp string) (string, error) {
	client := impl.getClient(P2pPeersDb)
	return client.Get(otp).Result()
}

func (impl *RegistryIml) GetNodeConfig() (string, error) {
	client := impl.getClient(P2pBootstrapNodeDb)
	return client.Get(BootstrapNode).Result()
}
