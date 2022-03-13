package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"jimmy_tech_crud_gin/config"
	"jimmy_tech_crud_gin/service/response/result"
)

var ErrNoAccess = errors.New("no access")

func NewTokenAuth(conf config.TokenAuthConf, serializer result.Setter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query(conf.QueryParam)

		if token == conf.Token {
			return
		}

		serializer.SetStatusError(ctx, http.StatusForbidden, ErrNoAccess)
		ctx.Abort()
	}
}
