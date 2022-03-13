package request

import "github.com/gin-gonic/gin"

type ConditionalBinder interface {
	Supports(ctx *gin.Context) bool
	Bind(ctx *gin.Context, obj interface{}) error
}
