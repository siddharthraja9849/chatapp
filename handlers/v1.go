package handlers

import (
	"chatapp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type V1 struct {
	Type string
}

func NewV1Handler() *V1{
	return &V1{Type: utils.V1Route}
}

func (v V1) CheckConnection(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": v.Type + " check connection route working"})
}