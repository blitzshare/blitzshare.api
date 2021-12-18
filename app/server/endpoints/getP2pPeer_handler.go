package endpoints

import (
	"net/http"

	"blitzshare.api/app/dependencies"
	"github.com/gin-gonic/gin"
)

func GetP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		otp := c.Params.ByName("otp")
		multiAddr, err := deps.Registry.Get(otp)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"multiAddr": multiAddr})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		}
	}
}
