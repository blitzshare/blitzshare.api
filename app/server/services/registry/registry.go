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

var registry map[string]string

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

type GetSet interface {
	set(d *deps.Dependencies, peer *model.PeerRegistry) bool
	getPeerRedisRecord(d *deps.Dependencies, pass string) (string, error)
}

func set(d *deps.Dependencies, key string, value string) bool {
	client := getClient(d)
	pong, err := client.Ping().Result()
	log.Infoln("set", pong, err)
	result, err := client.Set(key, value, time.Second*10000).Result()
	if err != nil {
		log.Errorf("client.Set err [%s]\n", err)
	} else {
		log.Infof("client.Set [%s]\n", result)
	}
	return err != nil
}

func get(d *deps.Dependencies, key string) (string, error) {
	client := getClient(d)
	result, err := client.Get(key).Result()
	if err != nil {
		log.Errorf("client.Get err [%s]\n", err)
	} else {
		log.Infof("client.Get [%s]\n", result)
	}
	return result, err
}

func RegisterPeer(d *deps.Dependencies, peer *model.PeerRegistry) (string, error) {
	if set(d, str.SanatizeStr(peer.OneTimePass), str.SanatizeStr(peer.MultiAddr)) {
		return peer.MultiAddr, nil
	}
	return "", nil
}

func GetPeerMultiAddr(d *deps.Dependencies, oneTimePass string) (string, error) {
	log.Infoln("GetPeerMultiAddr", oneTimePass)
	if registry != nil {
		result, err := get(d, oneTimePass)
		if err != nil {
			return result, err
		}
	}
	return "", errors.New("PeerNotFoundError")
}
