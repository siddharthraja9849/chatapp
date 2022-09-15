package handlers

import (
	"chatapp/types"
	"chatapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1 struct {
	Type string
}

func NewV1Handler() *V1 {
	return &V1{Type: utils.V1Route}
}

func (v V1) CheckConnection(ctx *gin.Context) {
	utils.BindResponse(ctx, http.StatusOK, map[string]string{"message": v.Type + " check connection route working"})
}

func (v V1) Register(ctx *gin.Context) {
	registerBody := types.ParseRegisterBody(ctx)
	utils.BindResponse(ctx, http.StatusOK, registerBody)
}
