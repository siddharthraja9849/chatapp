package handlers

import (
	"chatapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V2 struct {
	Type string
}

func NewV2Handler() *V2 {
	return &V2{
		Type: utils.V2Route,
	}
}

func (v V2) CheckConnection(ctx *gin.Context) {
	utils.BindResponse(ctx, http.StatusOK, map[string]string{"message": v.Type + " check connection route working"}, nil)
}
