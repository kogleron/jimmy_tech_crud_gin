package articles_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	pkgarticles "jimmy_tech_crud_gin/controllers/articles"
	mocks "jimmy_tech_crud_gin/mocks/service/storage"
	"jimmy_tech_crud_gin/models"
)

func TestList_Success(t *testing.T) {
	t.Parallel()

	articles := []models.Article{{}}
	filter := models.ArticlesFilter{
		OrderBy:  "id",
		OrderDir: "asc",
	}
	ctx := &gin.Context{
		Request: &http.Request{
			URL: &url.URL{},
		},
	}
	repository := &mocks.ArticlesRepository{}

	repository.
		On("FindByFilter", filter).Return(articles, nil)

	got, err := pkgarticles.List(ctx, repository)

	assert.Nil(t, err)
	assert.Equal(t, articles, got)
	repository.AssertExpectations(t)
}

func TestList_FindErr(t *testing.T) {
	t.Parallel()

	filter := models.ArticlesFilter{
		OrderBy:  "id",
		OrderDir: "asc",
	}
	ctx := &gin.Context{
		Request: &http.Request{
			URL: &url.URL{},
		},
	}
	repository := &mocks.ArticlesRepository{}

	repository.
		On("FindByFilter", filter).Return(nil, ErrDefault)

	got, err := pkgarticles.List(ctx, repository)

	assert.EqualError(t, err, "err")
	assert.Nil(t, got)
}
