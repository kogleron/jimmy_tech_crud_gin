package articles_test

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"jimmy_tech_crud_gin/controllers/articles"
	intrnlerr "jimmy_tech_crud_gin/internal/errors"
	mocks "jimmy_tech_crud_gin/mocks/service/storage"
	"jimmy_tech_crud_gin/models"
)

var ErrDefault = errors.New("err")

func TestGet_Success(t *testing.T) {
	t.Parallel()

	expectedArticle := models.Article{}
	ctx := &gin.Context{
		Params: gin.Params{
			gin.Param{Key: "id", Value: "5"},
		},
	}
	repository := &mocks.ArticlesRepository{}

	repository.
		On("FindByID", 5).Return(&expectedArticle, nil)

	got, err := articles.Get(ctx, repository)

	assert.Nil(t, err)
	assert.Equal(t, &expectedArticle, got)
	repository.AssertExpectations(t)
}

func TestGet_FindErr(t *testing.T) {
	t.Parallel()

	ctx := &gin.Context{
		Params: gin.Params{
			gin.Param{Key: "id", Value: "5"},
		},
	}
	repository := &mocks.ArticlesRepository{}

	repository.
		On("FindByID", mock.Anything).
		Return(nil, ErrDefault)

	got, err := articles.Get(ctx, repository)

	codeErr, ok := err.(intrnlerr.CodeError) //nolint:errorlint
	assert.True(t, ok)
	assert.Equal(t, codeErr.Code(), 404)
	assert.EqualError(t, err, "err")
	assert.Nil(t, got)
}

func TestGet_NoArticleId(t *testing.T) {
	t.Parallel()

	ctx := &gin.Context{}
	repository := &mocks.ArticlesRepository{}

	got, err := articles.Get(ctx, repository)

	codeErr, ok := err.(intrnlerr.CodeError) //nolint:errorlint
	assert.True(t, ok)
	assert.Equal(t, codeErr.Code(), 400)
	assert.EqualError(t, err, "strconv.ParseInt: parsing \"\": invalid syntax")
	assert.Nil(t, got)
}
