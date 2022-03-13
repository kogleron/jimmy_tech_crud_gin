package request

import (
	"github.com/gin-gonic/gin"
)

func NewJSONBinder() ConditionalBinder {
	return &jsonBinder{}
}

type jsonBinder struct{}

func (b *jsonBinder) Supports(ctx *gin.Context) bool {
	return ctx.GetHeader("Content-Type") == "application/json"
}

func (b *jsonBinder) Bind(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindJSON(obj)
}
