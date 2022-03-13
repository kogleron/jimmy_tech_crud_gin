package models_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"jimmy_tech_crud_gin/models"
)

func TestArticlesFilter_Errors(t *testing.T) {
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
		{
			name: "wrong orderDir",
			filter: models.ArticlesFilter{
				OrderBy:  "id",
				OrderDir: "bad\"'",
			},
			want: "Key: 'ArticlesFilter.OrderDir' Error:Field validation for 'OrderDir' failed on the 'oneof' tag",
		},
	}

	for n := range tests {
		tt := &tests[n]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			validatorInstance := validator.New()

			err := validatorInstance.Struct(tt.filter)

			assert.EqualError(t, err, tt.want)
		})
	}
}
