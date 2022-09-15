package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Payload interface{} `json:"data"`
}

func BindResponse(ctx *gin.Context, status int, data interface{}) {
	response := Response{Payload: data}
	ctx.JSON(status, response)
}
