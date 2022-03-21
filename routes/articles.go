package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jimmy_tech_crud_gin/controllers/articles"
	"jimmy_tech_crud_gin/service/request"
	"jimmy_tech_crud_gin/service/storage"
)

// NewArticlesListRoute godoc
// @Summary Retrieves articles
// @Param orderBy query string false "By which column to order articles" Enums(id, title, body)
// @Param orderDir query string false "Order direction" Enums(asc, desc)
// @Param offset query integer false "From which article to retrieve."
// @Param offset query integer false "From which article to retrieve."
// @Param limit query integer false "The number of articles to retrieve."
// @Produce json
// @Produce xml
// @Success 200 {array} models.Article
// @Router /articles [get]
func NewArticlesListRoute(repository storage.ArticlesRepository) Route {
	return &routeData{
		method: http.MethodGet,
		path:   "/articles",
		handler: func(ctx *gin.Context) (interface{}, error) {
			return articles.List(ctx, repository)
		},
	}
}

// NewArticlesUpdateRoute godoc
// @Summary Update an article
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Param token query string true "Access token (auth_token_xxx)"
// @Param id path integer true "Article ID"
// @Param article body models.UpdateArticleInput true "Data of an article"
// @Success 200 {object} models.Article
// @Router /articles/{id} [put]
func NewArticlesUpdateRoute(repository storage.ArticlesRepository, tokenAuth gin.HandlerFunc, binder request.Binder,
) Route {
	return &routeData{
		method:     http.MethodPut,
		path:       "/articles/:id",
		middleware: []gin.HandlerFunc{tokenAuth},
		handler: func(ctx *gin.Context) (interface{}, error) {
			return articles.Update(ctx, repository, binder)
		},
	}
}

// NewArticlesGetRoute godoc
// @Summary Retrieves an article
// @Description Allows to get an article in different formats (json, xml)
// @Produce json
// @Produce xml
// @Param id path integer true "Article ID"
// @Success 200 {object} models.Article
// @Router /articles/{id} [get]
func NewArticlesGetRoute(repository storage.ArticlesRepository) Route {
	return &routeData{
		method: http.MethodGet,
		path:   "/articles/:id",
		handler: func(ctx *gin.Context) (interface{}, error) {
			return articles.Get(ctx, repository)
		},
	}
}

// NewArticlesDeleteRoute godoc
// @Summary Deletes an article
// @Produce json
// @Produce xml
// @Param token query string true "Access token (auth_token_xxx)"
// @Param id path integer true "Article ID"
// @Success 200 {boolean} boolean
// @Router /articles/{id} [delete]
func NewArticlesDeleteRoute(repository storage.ArticlesRepository, tokenAuth gin.HandlerFunc) Route {
	return &routeData{
		method:     http.MethodDelete,
		path:       "/articles/:id",
		middleware: []gin.HandlerFunc{tokenAuth},
		handler: func(ctx *gin.Context) (interface{}, error) {
			return articles.Delete(ctx, repository)
		},
	}
}

// NewArticlesCreateRoute godoc
// @Summary Creates an article
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Param token query string true "Access token (auth_token_xxx)"
// @Param article body models.CreateArticleInput true "Data of an article"
// @Router /articles [post]
func NewArticlesCreateRoute(repository storage.ArticlesRepository, tokenAuth gin.HandlerFunc, binder request.Binder,
) Route {
	return &routeData{
		method:     http.MethodPost,
		path:       "/articles",
		middleware: []gin.HandlerFunc{tokenAuth},
		handler: func(ctx *gin.Context) (interface{}, error) {
			return articles.Create(ctx, repository, binder)
		},
	}
}
