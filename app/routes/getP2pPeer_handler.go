package routes

import (
	"blitzshare.api/app/model"
	"encoding/json"
	"net/http"

	"blitzshare.api/app/dependencies"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func parsePeerConfig(config string) (*model.P2pPeerRegistryResponse, error) {
	var conf model.P2pPeerRegistryResponse
	err := json.Unmarshal([]byte(config), &conf)
	log.Debugln("GetP2pPeerHandler", config)
	if err != nil {
		log.Errorln("GetP2pPeerHandler", err.Error())
	}
	return &conf, err
}

// GetP2pRegistry godoc
// @Summary  Returns peer registry config via OTP
// @Param        X-Api-Key  header  string  true  "API key authenication header"
// @Schemes
// @Tags    p2p-registry
// @Param    otp    query     string  false  "blitzshare generated otp"
// @Success 200 {object} model.P2pPeerRegistryResponse "peer config used for libp2p connection"
// @Success 404 "otp not found"
// @Success 500 "internal server error"
// @Router   /p2p/registry/:otp [get]
func GetP2pRegistryHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		if IsNotAuthorized(c, deps.ApiKeychain) {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}
		otp := c.Params.ByName("otp")
		log.Debugln("GetP2pPeerHandler", otp)
		str, err := deps.Registry.GetOtp(otp)
		resp, parseErr := parsePeerConfig(str)
		if err == nil && parseErr == nil {
			c.JSON(http.StatusOK, resp)
		} else {
			if err != nil {
				log.Errorln("GetP2pPeerHandler", err.Error())
				c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
			} else if parseErr != nil {
				log.Errorln("GetP2pPeerHandler", parseErr.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"err": parseErr.Error()})
			}
		}
	}
}
