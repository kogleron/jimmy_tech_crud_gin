package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jimmy_tech_crud_gin/internal/errors"
	"jimmy_tech_crud_gin/models"
	"jimmy_tech_crud_gin/service/request"
	"jimmy_tech_crud_gin/service/storage"
)

func Update(ctx *gin.Context, repository storage.ArticlesRepository, binder request.Binder) (*models.Article, error) {
	var articleInput models.UpdateArticleInput

	err := binder.Bind(ctx, &articleInput)
	if nil != err {
		return nil, errors.WithCode(err, http.StatusBadRequest)
	}

	article, err := Get(ctx, repository)
	if err != nil {
		return nil, err
	}

	if nil == article {
		return nil, ErrNotFound
	}

	article = repository.Update(article, articleInput)

	return article, nil
}
