package result

import (
	"github.com/gin-gonic/gin"
)

func NewXMLSetter() ConditionalSetter {
	return &xmlSetter{}
}

type xmlSetter struct{}

func (s *xmlSetter) Supports(ctx *gin.Context) bool {
	contentType := ctx.GetHeader("Accept")
	switch contentType {
	case "text/xml", "application/xml":
		return true
	default:
		return false
	}
}

func (s *xmlSetter) Set(ctx *gin.Context, status int, data interface{}) {
	ctx.XML(status, data)
}
