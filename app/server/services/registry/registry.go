package registry

import (
	"errors"
	"time"

	deps "blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server/services/str"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type PeerNotFoundError struct {
	arg  int
	prob string
}

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

func set(d *deps.Dependencies, key string, value string) bool {
	client := getClient(d)
	_, err := client.Set(key, value, time.Second*10000).Result()
	if err != nil {
		log.Errorf("client.Set err", err)
		return false
	} else {
		log.Infof("client.Set {%s : %s}", key, value)
		return true
	}
}

func get(d *deps.Dependencies, key string) (string, error) {
	client := getClient(d)
	return client.Get(key).Result()
}

func RegisterPeer(d *deps.Dependencies, peer *model.P2pPeerRegistryCmd) error {
	if set(d, str.SanatizeStr(peer.OneTimePass), str.SanatizeStr(peer.MultiAddr)) {
		return nil
	}
	return errors.New("Failed to register peer")
}

func GetPeerMultiAddr(d *deps.Dependencies, oneTimePass string) (string, error) {
	log.Infoln("GetPeerMultiAddr", oneTimePass)
	result, err := get(d, oneTimePass)
	if err == nil {
		return result, err
	}
	return "", errors.New("PeerNotFoundError")
}
