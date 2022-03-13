package result

import (
	"github.com/gin-gonic/gin"
)

func NewJSONSetter() ConditionalSetter {
	return &jsonSetter{}
}

type jsonSetter struct{}

func (s *jsonSetter) Supports(ctx *gin.Context) bool {
	return ctx.GetHeader("Accept") == "application/json"
}

func (s *jsonSetter) Set(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}
