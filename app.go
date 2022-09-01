package main

import (
	"chatapp/types"
	"chatapp/utils"
	"fmt"
)

func main() {
	config := types.ENV{}
	utils.LoadEnvConfigs(&config)
	fmt.Println(config.TestDbUsername)
}
