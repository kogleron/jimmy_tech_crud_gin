package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jimmy_tech_crud_gin/internal/errors"
	"jimmy_tech_crud_gin/models"
	"jimmy_tech_crud_gin/service/request"
	"jimmy_tech_crud_gin/service/storage"
)

func Create(ctx *gin.Context, repository storage.ArticlesRepository, binder request.Binder) (*models.Article, error) {
	var articleInput models.CreateArticleInput

	err := binder.Bind(ctx, &articleInput)
	if nil != err {
		return nil, errors.WithCode(err, http.StatusBadRequest)
	}

	article := &models.Article{
		Title: articleInput.Title,
		Body:  articleInput.Body,
	}

	article = repository.Create(article)

	return article, nil
}
