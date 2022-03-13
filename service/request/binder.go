package request

import "github.com/gin-gonic/gin"

type Binder interface {
	Bind(ctx *gin.Context, obj interface{}) error
}
