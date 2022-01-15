package server

import (
	"net/http"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/server/endpoints"
	"github.com/gin-gonic/gin"
)

func registerDefaultRoute(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		endpoints.AddDefaultResponseHeaders(c)
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found :("})
	})
}
func NewRouter(deps *dependencies.Dependencies) *gin.Engine {
	router := gin.New()
	if deps.Config.Settings.Environment != "local" {
		gin.SetMode(gin.ReleaseMode)
	}
	registerDefaultRoute(router)

	router.GET("/test", endpoints.HealthHandler())
	// p2p/registy
	router.POST("/p2p/registry", endpoints.RegisterP2pPeerHandler(deps))
	router.GET("/p2p/registry/:otp", endpoints.GetP2pPeerHandler(deps))
	router.DELETE("/p2p/registry/:otp/:token", endpoints.DeleteP2pPeerHandler(deps))
	// p2p/boostrap-node
	router.GET("/p2p/bootstrap-node", endpoints.GetP2pBootstrapNode(deps))
	return router
}
