package endpoints

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"blitzshare.api/app/dependencies"
	"github.com/gin-gonic/gin"
)

func GetP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		log.Infoln("GetP2pPeerHandler id:", id)
		multiAddr, err := deps.Registry.Get(id)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"multiAddr": multiAddr})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		}
	}
}
