package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Payload interface{} `json:"data"`
	Error   interface{} `json:"errorMessage,omitempty"`
}

func BindResponse(ctx *gin.Context, status int, data interface{}, err error) {
	var response Response
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Payload = data
	}
	ctx.JSON(status, response)
}
