package models

type ArticlesFilter struct {
	// Column
	OrderBy string `form:"orderBy" binding:"oneof=id title body" validate:"oneof=id title body"`
	// Order direction: asc|desc
	OrderDir string `form:"orderDir" binding:"oneof=asc desc" validate:"oneof=asc desc"`
	Offset   int    `form:"offset" binding:"numeric"`
	Limit    int    `form:"limit" binding:"numeric"`
}
