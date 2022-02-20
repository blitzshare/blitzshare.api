package routes

import (
	"blitzshare.api/app/model"
	"encoding/json"
	"net/http"

	"blitzshare.api/app/dependencies"
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

// BootstrapNode godoc
// @Summary  Returns bootstrap node configuration for for libp2p p2p connection
// @Param        X-Api-Key  header  string  true  "API key authenication header"
// @Schemes
// @Tags    bootstrap-node
// @Produce json
// @Success 200 {object} model.NodeConfig "bootstrap node config"
// @Success 500 "failed to fetch node config"
// @Router   /p2p/bootstrap-node [get]
func GetBootstrapNodeHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		if IsNotAuthorized(c, deps.ApiKeychain) {
			c.JSON(http.StatusUnauthorized, "")
			return
		}
		config, err := deps.Registry.GetNodeConfig()
		if err != nil || config == "" {
			log.Errorln("cannot obtain node config", config, err)
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
