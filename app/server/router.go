package server

import (
	"github.com/blitzshare/blitzshare.fileshare.api/app/dependencies"
	"github.com/blitzshare/blitzshare.fileshare.api/app/server/endpoints"
	"github.com/gin-gonic/gin"
)

func NewRouter(deps *dependencies.Dependencies) *gin.Engine {
	router := gin.New()
	router.GET("/test", endpoints.HealthHandler())
	router.GET("/file-share", endpoints.FileShareHandler(deps))

	return router
}
