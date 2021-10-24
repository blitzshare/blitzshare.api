package server

import (
	"blitzshare.fileshare.api/app/dependencies"
	"blitzshare.fileshare.api/app/server/endpoints"
	"github.com/gin-gonic/gin"
)

func NewRouter(deps *dependencies.Dependencies) *gin.Engine {
	router := gin.New()
	if deps.Config.Settings.Environment != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.GET("/test", endpoints.HealthHandler())
	router.GET("/file-share", endpoints.FileShareHandler(deps))

	return router
}
