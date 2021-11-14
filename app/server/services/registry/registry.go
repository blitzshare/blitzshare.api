package registry

import (
	deps "blitzshare.api/app/dependencies"
	"errors"

	"blitzshare.api/app/model"
	log "github.com/sirupsen/logrus"
)

var registry map[string]string

type PeerNotFoundError struct {
	arg  int
	prob string
}

func RegisterPeer(d *deps.Dependencies, peer *model.PeerRegistry) string {
	log.Infoln("RegisterPeer", peer)
	if registry == nil {
		registry = make(map[string]string)
	}
	registry[peer.OneTimePass] = peer.MultiAddr
	return registry[peer.OneTimePass]
}

func GetPeerMultiAddr(d *deps.Dependencies, oneTimePass string) (string, error) {
	log.Infoln("GetPeerMultiAddr", oneTimePass)
	if registry != nil {
		if multiAddress, ok := registry[oneTimePass]; ok {
			return multiAddress, nil
		}
	}
	return "", errors.New("PeerNotFoundError")
}
