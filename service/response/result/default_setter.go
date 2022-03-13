package result

import (
	"github.com/gin-gonic/gin"
)

func NewDefaultSetter(setter ConditionalSetter) ConditionalSetter {
	return &defaultSetter{setter: setter}
}

type defaultSetter struct {
	setter ConditionalSetter
}

func (s *defaultSetter) Supports(_ *gin.Context) bool {
	return true
}

func (s *defaultSetter) Set(ctx *gin.Context, status int, data interface{}) {
	s.setter.Set(ctx, status, data)
}
