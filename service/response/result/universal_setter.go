package result

import (
	"net/http"

	"github.com/gin-gonic/gin"

	intrnlerr "jimmy_tech_crud_gin/internal/errors"
)

func NewUniversalSetter(setters []ConditionalSetter) Setter {
	return &universalSetter{
		setters: setters,
	}
}

type universalSetter struct {
	setters []ConditionalSetter
}

func (s *universalSetter) SetStatusError(ctx *gin.Context, status int, err error) {
	s.Set(
		ctx,
		status,
		gin.H{"error": err.Error()},
	)
}

func (s *universalSetter) Set(ctx *gin.Context, status int, data interface{}) {
	for _, setter := range s.setters {
		if setter.Supports(ctx) {
			setter.Set(ctx, status, data)

			break
		}
	}
}

func (s *universalSetter) SetError(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError

	codeErr, ok := err.(intrnlerr.CodeError) //nolint:errorlint
	if ok {
		status = codeErr.Code()
	}

	s.Set(ctx, status, gin.H{"error": err.Error()})
}

func (s *universalSetter) SetEmpty(ctx *gin.Context, status int) {
	ctx.Status(status)
}

func (s *universalSetter) SetSuccess(ctx *gin.Context, data interface{}) {
	s.Set(
		ctx,
		http.StatusOK,
		gin.H{"data": data},
	)
}
