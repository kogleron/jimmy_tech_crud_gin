package result

import "github.com/gin-gonic/gin"

type ConditionalSetter interface {
	Supports(ctx *gin.Context) bool
	Set(ctx *gin.Context, status int, data interface{})
}
