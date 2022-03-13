package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"jimmy_tech_crud_gin/config"
	// Need for swagger docs
	_ "jimmy_tech_crud_gin/docs"
)

func NewSwaggerRoute(conf *config.ServerConf) Route {
	swaggerURL := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.Port))

	return &routeData{
		method: http.MethodGet,
		path:   "/swagger/*any",
		middleware: []gin.HandlerFunc{
			ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL),
		},
	}
}
