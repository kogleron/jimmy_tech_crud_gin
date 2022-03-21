package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"jimmy_tech_crud_gin/config"
	"jimmy_tech_crud_gin/middleware"
	"jimmy_tech_crud_gin/routes"
	"jimmy_tech_crud_gin/service"
	"jimmy_tech_crud_gin/service/request"
	"jimmy_tech_crud_gin/service/response/result"
	"jimmy_tech_crud_gin/service/storage"
)

// @title Yet another CRUD service
// @version 1.0
// @description This is a test for the job.
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @schemes http
func main() { //nolint:funlen
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}

	serverConf, err := config.GetServerConf()
	if err != nil {
		log.Panic(err)
	}

	tokenAuthConf := config.GetTokenAuthConf()

	dbConf, err := config.GetDBConf()
	if err != nil {
		log.Panic(err)
	}

	db, err := service.ConnectDB(*dbConf)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		err = db.Close()
		if nil != err {
			log.Panic(err)
		}
	}()

	resultSerializer := result.NewUniversalSetter([]result.ConditionalSetter{
		result.NewXMLSetter(),
		result.NewDefaultSetter(result.NewJSONSetter()),
	})
	binder := request.NewUniversalBinder([]request.ConditionalBinder{
		request.NewJSONBinder(),
		request.NewXMLBinder(),
	})

	articlesRepository := storage.NewArticlesRepository(db)

	// Middleware
	tokenAuth := middleware.NewTokenAuth(tokenAuthConf, resultSerializer)

	engine := gin.Default()
	// Routes
	routesData := []routes.Route{
		routes.NewArticlesListRoute(articlesRepository),
		routes.NewArticlesGetRoute(articlesRepository),
		routes.NewArticlesCreateRoute(articlesRepository, tokenAuth, binder),
		routes.NewArticlesUpdateRoute(articlesRepository, tokenAuth, binder),
		routes.NewArticlesDeleteRoute(articlesRepository, tokenAuth),
		routes.NewSwaggerRoute(serverConf),
	}
	for i := range routesData {
		route := routesData[i]
		handlers := route.GetMiddleware()

		if nil != route.GetHandler() {
			handlers = append(handlers, func(ctx *gin.Context) {
				data, err := route.GetHandler()(ctx)
				if err != nil {
					resultSerializer.SetError(ctx, err)

					return
				}

				resultSerializer.SetSuccess(ctx, data)
			})
		}

		engine.Handle(route.GetMethod(), route.GetPath(), handlers...)
	}

	err = engine.Run(":" + strconv.FormatInt(serverConf.Port, 10)) //nolint:gomnd
	if err != nil {
		panic(err)
	}
}
