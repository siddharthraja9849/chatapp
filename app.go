package main

import (
	db "chatapp/db"
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

	psql := db.NewPsqlConfig(config.AppDbHost, config.AppDbPort, config.AppDbUsername, config.AppDbPassword, config.AppDbName, config.AppDbSslMode)

	log.Println(psql)

	router := gin.Default()

	router.Use(gin.Logger())

	routes.RouterGroups(utils.V1Route, router, psql)
	routes.RouterGroups(utils.V2Route, router, psql)

	if err := router.Run(":" + config.PORT); err != nil {
		log.Printf("%v+", err)
	}

}
