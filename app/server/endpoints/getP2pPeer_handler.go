package endpoints

import (
	"encoding/json"
	"net/http"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func parsePeerConfig(config string) (*model.P2pPeerRegistryCmd, error) {
	var conf model.P2pPeerRegistryCmd
	err := json.Unmarshal([]byte(config), &conf)
	if err != nil {
		log.Errorln("GetP2pPeerHandler", err.Error())
	}
	return &conf, err
}
func GetP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		otp := c.Params.ByName("otp")
		log.Debugln("GetP2pPeerHandler", otp)
		str, err := deps.Registry.GetOtp(otp)
		nodeConfig, parseErr := parsePeerConfig(str)
		if err == nil && parseErr == nil {
			c.JSON(http.StatusOK, nodeConfig)
		} else {
			if err != nil {
				log.Errorln("GetP2pPeerHandler", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			} else if parseErr != nil {
				log.Errorln("GetP2pPeerHandler", parseErr.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"err": parseErr.Error()})
			}
		}
	}
}
