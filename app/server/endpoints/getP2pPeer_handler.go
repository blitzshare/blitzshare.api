package endpoints

import (
	"net/http"

	"blitzshare.api/app/dependencies"
	"github.com/gin-gonic/gin"
)

func GetP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		id := c.Params.ByName("id")
		multiAddr, err := deps.Registry.Get(id)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"multiAddr": multiAddr})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		}
	}
}
