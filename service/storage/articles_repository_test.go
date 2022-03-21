package storage_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"jimmy_tech_crud_gin/models"
	"jimmy_tech_crud_gin/service/storage"
)

func Test_ImplArticlesRepository_FindByFilter_Validation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		filter models.ArticlesFilter
		want   string
	}{
		{
			name: "wrong orderBy",
			filter: models.ArticlesFilter{
				OrderBy:  "bad\"'",
				OrderDir: "asc",
			},
			want: "Key: 'ArticlesFilter.OrderBy' Error:Field validation for 'OrderBy' failed on the 'oneof' tag",
		},
	}
	for n := range tests {
		tt := &tests[n]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := storage.NewArticlesRepository(&gorm.DB{})

			got, err := repository.FindByFilter(tt.filter)

			assert.Nil(t, got)
			assert.EqualError(t, err, tt.want)
		})
	}
}
