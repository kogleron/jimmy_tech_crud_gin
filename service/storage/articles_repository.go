package storage

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"jimmy_tech_crud_gin/models"
)

type ArticlesRepository interface {
	FindByID(id int) (*models.Article, error)
	Delete(article *models.Article)
	Create(article *models.Article) *models.Article
	Update(article *models.Article, data interface{}) *models.Article
	FindByFilter(filter models.ArticlesFilter) ([]models.Article, error)
}

type ImplArticlesRepository struct {
	db *gorm.DB
}

func (r *ImplArticlesRepository) FindByFilter(filter models.ArticlesFilter) ([]models.Article, error) {
	validatorInstance := validator.New()

	err := validatorInstance.Struct(filter)
	if err != nil {
		return nil, err
	}

	db := r.db.
		Order(filter.OrderBy + " " + filter.OrderDir).
		Offset(filter.Offset)

	if filter.Limit > 0 {
		db = db.Limit(filter.Limit)
	}

	var articles []models.Article

	err = db.Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ImplArticlesRepository) Update(article *models.Article, data interface{}) *models.Article {
	r.db.Model(&article).Updates(data)

	return article
}

func (r *ImplArticlesRepository) Delete(article *models.Article) {
	r.db.Delete(article)
}

func (r *ImplArticlesRepository) Create(article *models.Article) *models.Article {
	r.db.Create(article)

	return article
}

func (r *ImplArticlesRepository) FindByID(id int) (*models.Article, error) {
	var article models.Article

	err := r.db.
		Where("id = ?", id).
		First(&article).
		Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func NewArticlesRepository(db *gorm.DB) ArticlesRepository {
	return &ImplArticlesRepository{
		db: db,
	}
}
