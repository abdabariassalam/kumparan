package repository

import (
	"github.com/bariasabda/kumparan/domain/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	GetArticle(search string, author string) (*[]entity.Article, error)
	CreateArticle(article entity.Article) (*entity.Article, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}
