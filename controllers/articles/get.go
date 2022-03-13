package articles

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jimmy_tech_crud_gin/internal/errors"
	"jimmy_tech_crud_gin/models"
	"jimmy_tech_crud_gin/service/storage"
)

func Get(ctx *gin.Context, repository storage.ArticlesRepository) (*models.Article, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64) //nolint:gomnd
	if err != nil {
		return nil, errors.WithCode(err, http.StatusBadRequest)
	}

	article, err := repository.FindByID(int(id))
	if err != nil {
		return nil, errors.WithCode(err, http.StatusNotFound)
	}

	return article, nil
}
