package handlers

import (
	"chatapp/types"
	"chatapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type V1 struct {
	Type string
	Psql *gorm.DB
}

func NewV1Handler(psql *gorm.DB) *V1 {
	return &V1{Type: utils.V1Route, Psql: psql}
}

func (v V1) CheckConnection(ctx *gin.Context) {
	utils.BindResponse(ctx, http.StatusOK, map[string]string{"message": v.Type + " check connection route working"}, nil)
}

func (v V1) Register(ctx *gin.Context) {
	var registerBody types.RegisterBody
	var err error
	if registerBody, err = types.ParseRegisterBody(ctx); err != nil {
		utils.BindResponse(ctx, http.StatusBadRequest, registerBody, err)
		return
	}
	utils.BindResponse(ctx, http.StatusOK, registerBody, err)
}
