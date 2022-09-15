package handlers

import (
	"chatapp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type V2 struct{
	Type string
}

func NewV2Handler() (*V2){
	return &V2{
		Type: utils.V2Route,
	}
}


func (v V2) CheckConnection(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": v.Type + " check connection route working"})
}
