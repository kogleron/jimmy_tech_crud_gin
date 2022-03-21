package articles_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"jimmy_tech_crud_gin/controllers/articles"
	intrnlerr "jimmy_tech_crud_gin/internal/errors"
	mocks "jimmy_tech_crud_gin/mocks/service/storage"
	"jimmy_tech_crud_gin/models"
)

func TestDelete_Success(t *testing.T) {
	t.Parallel()

	article := &models.Article{}
	ctx := &gin.Context{
		Params: gin.Params{
			gin.Param{Key: "id", Value: "5"},
		},
	}
	repository := &mocks.ArticlesRepository{}

	repository.
		On("FindByID", 5).
		Return(article, nil)
	repository.
		On("Delete", article)

	got, err := articles.Delete(ctx, repository)

	assert.Nil(t, err)
	assert.True(t, got)
	repository.AssertExpectations(t)
}

func TestDelete_Err(t *testing.T) {
	t.Parallel()

	article := &models.Article{}
	ctx := &gin.Context{
		Params: gin.Params{
			gin.Param{Key: "id", Value: "5"},
		},
	}
	repository := &mocks.ArticlesRepository{}

	repository.
		On("FindByID", 5).
		Return(nil, ErrDefault)
	repository.
		On("Delete", article)

	got, err := articles.Delete(ctx, repository)

	codeErr, ok := err.(intrnlerr.CodeError) //nolint:errorlint
	assert.True(t, ok)
	assert.Equal(t, codeErr.Code(), 404)
	assert.EqualError(t, err, "err")
	assert.False(t, got)
}
