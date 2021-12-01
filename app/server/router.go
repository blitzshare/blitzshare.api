package server

import (
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/server/endpoints"
	"github.com/gin-gonic/gin"
)

func NewRouter(deps *dependencies.Dependencies) *gin.Engine {
	router := gin.New()
	if deps.Config.Settings.Environment != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.GET("/test", endpoints.HealthHandler())
	//router.GET("/file-share-link", endpoints.FileShareHandler(deps))
	router.POST("/p2p/registry", endpoints.RegisterP2pPeerHandler(deps))
	router.GET("/p2p/registry/:id", endpoints.GetP2pPeerHandler(deps))

	return router
}
