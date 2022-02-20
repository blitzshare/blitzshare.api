package registry

import (
	"github.com/go-redis/redis"
)

var p2pPeersDbClient *redis.Client
var p2pBootstraoNodeDbClient *redis.Client

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

func New(redisUrl string) Registry {
	return &RegistryIml{redisUrl: redisUrl}
}

func (r *RegistryIml) getPeersClient() *redis.Client {
	if p2pPeersDbClient == nil {
		p2pPeersDbClient = redis.NewClient(&redis.Options{
			Addr:     r.redisUrl,
			Password: "",
			DB:       P2pPeersDb,
		})
	}
	return p2pPeersDbClient
}
func (r *RegistryIml) getBootstrapNodeDbClient() *redis.Client {
	if p2pBootstraoNodeDbClient == nil {
		p2pBootstraoNodeDbClient = redis.NewClient(&redis.Options{
			Addr:     r.redisUrl,
			Password: "",
			DB:       P2pBootstrapNodeDb,
		})
	}
	return p2pBootstraoNodeDbClient
}

func (impl *RegistryIml) GetOtp(otp string) (string, error) {
	client := impl.getPeersClient()
	return client.Get(otp).Result()
}

func (impl *RegistryIml) GetNodeConfig() (string, error) {
	client := impl.getBootstrapNodeDbClient()
	return client.Get(BootstrapNode).Result()
}
