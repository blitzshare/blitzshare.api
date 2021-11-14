package registry

import (
	"errors"
	"time"

	deps "blitzshare.api/app/dependencies"

	"blitzshare.api/app/model"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var registry map[string]string

type PeerNotFoundError struct {
	arg  int
	prob string
}

var client *redis.Client

func getClient(d *deps.Dependencies) *redis.Client{
	if client == nil{
		client = redis.NewClient(&redis.Options{
			Addr:     d.Config.Settings.RedisUrl,
			Password: "",
			DB:       0,
		})
	}
	return client
}

func setPeerRedisRecord(d *deps.Dependencies, peer *model.PeerRegistry) {
	client := getClient(d)
	pong, err := client.Ping().Result()
	log.Infoln("setPeerRedisRecord", pong, err)
	err = client.Set(peer.OneTimePass, peer.MultiAddr, time.Second).Err()
	log.Infoln("client.Set", err)
}


func getPeerRedisRecord(d *deps.Dependencies, pass string ) string{
	client := getClient(d)
	log.Infoln("getPeerRedisRecord", pass)
	result, err := client.Get(pass).Result()
	if err != nil {
		log.Errorln(err)
	}
	log.Infoln("client.Get", result)
	return  result
}

func RegisterPeer(d *deps.Dependencies, peer *model.PeerRegistry) string {
	log.Infoln("RegisterPeer", peer)
	if registry == nil {
		registry = make(map[string]string)
	}
	registry[peer.OneTimePass] = peer.MultiAddr
	setPeerRedisRecord(d, peer)
	return registry[peer.OneTimePass]
}

func GetPeerMultiAddr(d *deps.Dependencies, oneTimePass string) (string, error) {
	log.Infoln("GetPeerMultiAddr", oneTimePass)
	if registry != nil {
		getPeerRedisRecord(d, oneTimePass)
		if multiAddress, ok := registry[oneTimePass]; ok {
			return multiAddress, nil
		}
	}
	return "", errors.New("PeerNotFoundError")
}
