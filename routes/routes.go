package routes

import (
	"chatapp/handlers"
	"chatapp/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouterGroups(version string, router *gin.Engine, psql *gorm.DB) {
	versionGroup := router.Group(version)
	switch version {
	case utils.V1Route:
		v1 := handlers.NewV1Handler(psql)
		v1Routes(versionGroup, v1)
	case utils.V2Route:
		v2 := handlers.NewV2Handler()
		v2Routes(versionGroup, v2)
	}
}

func v1Routes(versionGroup *gin.RouterGroup, v1 *handlers.V1) {
	versionGroup.GET(utils.CheckConnRoute, v1.CheckConnection)
	versionGroup.POST(utils.Register, v1.Register)
}

func v2Routes(versionGroup *gin.RouterGroup, v2 *handlers.V2) {
	versionGroup.GET(utils.CheckConnRoute, v2.CheckConnection)
}
