package routes

import (
	"net/http"

	"blitzshare.api/app/dependencies"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func registerDefaultRoute(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found :("})
	})
}

// @title Blitzshare API
// @version 1.0
func NewRouter(deps *dependencies.Dependencies) *gin.Engine {
	router := gin.New()
	if deps.Config.Settings.Environment != "local" {
		gin.SetMode(gin.ReleaseMode)
	}
	registerDefaultRoute(router)

	router.GET("/test", HealthHandler())
	// p2p/registy
	router.POST("/p2p/registry", PostP2pRegistryHandler(deps))
	router.GET("/p2p/registry/:otp", GetP2pRegistryHandler(deps))
	router.DELETE("/p2p/registry/:otp/:token", DeleteP2pPeerHandler(deps))
	// p2p/boostrap-node
	router.GET("/p2p/bootstrap-node", GetBootstrapNodeHandler(deps))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
