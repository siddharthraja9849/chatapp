package interfaces

import "github.com/gin-gonic/gin"

type Routes interface {
	CheckConnection(ctx *gin.Context)
}