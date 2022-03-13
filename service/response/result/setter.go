package result

import (
	"github.com/gin-gonic/gin"
)

type Setter interface {
	Set(ctx *gin.Context, status int, data interface{})
	SetError(ctx *gin.Context, err error)
	SetEmpty(ctx *gin.Context, status int)
	SetSuccess(ctx *gin.Context, data interface{})
	SetStatusError(ctx *gin.Context, status int, err error)
}
