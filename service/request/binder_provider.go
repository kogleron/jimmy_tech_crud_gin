package request

import (
	"github.com/gin-gonic/gin"
)

func NewUniversalBinder(binders []ConditionalBinder) Binder {
	return &universalBinder{
		binders: binders,
	}
}

type universalBinder struct {
	binders []ConditionalBinder
}

func (b *universalBinder) Bind(ctx *gin.Context, obj interface{}) error {
	for _, binder := range b.binders {
		if binder.Supports(ctx) {
			return binder.Bind(ctx, obj)
		}
	}

	return ErrNoBinder
}
