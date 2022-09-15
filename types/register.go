package types

import (
	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

func ParseRegisterBody(ctx *gin.Context) (RegisterBody, error) {
	var register RegisterBody
	if err := ctx.BindJSON(&register); err != nil {
		return register, err
	}
	return register, nil
}
