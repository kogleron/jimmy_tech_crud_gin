package articles_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"jimmy_tech_crud_gin/controllers/articles"
	intrnlerr "jimmy_tech_crud_gin/internal/errors"
	reqmocks "jimmy_tech_crud_gin/mocks/service/request"
	storagemocks "jimmy_tech_crud_gin/mocks/service/storage"
	"jimmy_tech_crud_gin/models"
)

func TestCreate_Success(t *testing.T) {
	t.Parallel()

	article := models.Article{Title: "t", Body: "b"}
	ctx := &gin.Context{}
	binder := &reqmocks.Binder{}
	repository := &storagemocks.ArticlesRepository{}

	binder.
		On("Bind", ctx, &models.CreateArticleInput{}).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.CreateArticleInput) //nolint:forcetypeassert
			arg.Title = "t"
			arg.Body = "b"
		}).
		Return(nil)
	repository.
		On("Create", &article).Return(&article)

	got, err := articles.Create(ctx, repository, binder)

	assert.Nil(t, err)
	assert.Equal(t, &article, got)
	binder.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestCreate_BindErr(t *testing.T) {
	t.Parallel()

	ctx := &gin.Context{}
	binder := &reqmocks.Binder{}
	repository := &storagemocks.ArticlesRepository{}

	binder.
		On("Bind", ctx, &models.CreateArticleInput{}).
		Return(ErrDefault)

	got, err := articles.Create(ctx, repository, binder)

	codeErr, ok := err.(intrnlerr.CodeError) //nolint:errorlint
	assert.True(t, ok)
	assert.Equal(t, codeErr.Code(), 400)
	assert.EqualError(t, err, "err")
	assert.Nil(t, got)
}
