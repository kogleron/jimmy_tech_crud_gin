package request

import (
	"github.com/gin-gonic/gin"
)

func NewXMLBinder() ConditionalBinder {
	return &xmlBinder{}
}

type xmlBinder struct{}

func (b *xmlBinder) Supports(ctx *gin.Context) bool {
	contentType := ctx.GetHeader("Content-Type")
	switch contentType {
	case "text/xml", "application/xml":
		return true
	default:
		return false
	}
}

func (b *xmlBinder) Bind(ctx *gin.Context, obj interface{}) error {
	return ctx.ShouldBindXML(obj)
}
