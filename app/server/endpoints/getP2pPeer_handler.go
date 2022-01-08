package endpoints

import (
	"net/http"

	"blitzshare.api/app/dependencies"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		otp := c.Params.ByName("otp")
		log.Infoln("GetP2pPeerHandler", otp)
		multiAddr, err := deps.Registry.Get(otp)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"multiAddr": multiAddr})
		} else {
			log.Errorln("GetP2pPeerHandler", err.Error())
			c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		}
	}
}
