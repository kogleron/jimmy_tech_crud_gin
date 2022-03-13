package models

type CreateArticleInput struct {
	Title string `json:"title" xml:"title" binding:"required,max=255"`
	Body  string `json:"body" xml:"body" binding:"required,max=65535"`
}

type UpdateArticleInput struct {
	Title string `json:"title" xml:"title" binding:"max=255"`
	Body  string `json:"body" xml:"body" binding:"max=65535"`
}
