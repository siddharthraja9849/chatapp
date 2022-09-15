package routes

import (
	"chatapp/handlers"
	"chatapp/utils"
	"github.com/gin-gonic/gin"
)


func RouterGroups(version string, router *gin.Engine) {
	versionGroup := router.Group(version)
	switch version {
	case utils.V1Route:
		v1 := handlers.NewV1Handler()
		versionGroup.GET(utils.CheckConnRoute, v1.CheckConnection)
	case utils.V2Route:
		v2 := handlers.NewV2Handler()
		versionGroup.GET(utils.CheckConnRoute, v2.CheckConnection)
	}
}
