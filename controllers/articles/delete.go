package articles

import (
	"github.com/gin-gonic/gin"
	"jimmy_tech_crud_gin/service/storage"
)

func Delete(ctx *gin.Context, repository storage.ArticlesRepository) (bool, error) {
	article, err := Get(ctx, repository)
	if err != nil {
		return false, err
	}

	repository.Delete(article)

	return true, nil
}
