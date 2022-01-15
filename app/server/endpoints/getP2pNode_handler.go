package endpoints

import (
	"encoding/json"
	"net/http"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func parseNodeConfig(config string) (*model.NodeConfig, error) {
	var nodeConfig model.NodeConfig
	err := json.Unmarshal([]byte(config), &nodeConfig)
	if err != nil {
		log.Errorln("GetP2pPeerHandler", err.Error())
	}
	return &nodeConfig, err
}

func GetP2pBootstrapNode(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		log.Debugln("GetP2pBootstrapNode")
		config, err := deps.Registry.GetNodeConfig()
		if err != nil || config == "" {
			c.JSON(http.StatusNotFound, "")
			return
		} else {
			nodeConfig, parseErr := parseNodeConfig(config)
			if parseErr == nil {
				log.Debugln("GetP2pBootstrapNode", nodeConfig)
				c.JSON(http.StatusOK, nodeConfig)
			} else {
				log.Errorln("GetP2pPeerHandler", parseErr.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"err": parseErr.Error()})
			}
		}
	}
}
