package main

import (
	"chatapp/routes"
	"chatapp/types"
	"chatapp/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.ForceConsoleColor()

	config := types.ENV{}
	utils.LoadEnvConfigs(&config)

	router := gin.Default()

	router.Use(gin.Logger())

	routes.RouterGroups(utils.V1Route, router)
	routes.RouterGroups(utils.V2Route, router)

	if err := router.Run(":" + config.PORT); err != nil {
		log.Printf("%v+", err)
	}

}
