package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jimmy_tech_crud_gin/internal/errors"
	"jimmy_tech_crud_gin/models"
	"jimmy_tech_crud_gin/service/storage"
)

func List(ctx *gin.Context, repository storage.ArticlesRepository) ([]models.Article, error) {
	filter := models.ArticlesFilter{
		OrderBy:  "id",
		OrderDir: "asc",
	}

	err := ctx.ShouldBindQuery(&filter)
	if nil != err {
		return nil, errors.WithCode(err, http.StatusBadRequest)
	}

	articles, err := repository.FindByFilter(filter)
	if nil != err {
		return nil, err
	}

	return articles, nil
}
