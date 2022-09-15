package types

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Username string `json:"username"`
}

func ParseRegisterBody(ctx *gin.Context) RegisterBody {
	var register RegisterBody
	if err := ctx.BindJSON(&register); err != nil {
		fmt.Errorf("%v+", err)
		os.Exit(1)
	}
	return register
}
