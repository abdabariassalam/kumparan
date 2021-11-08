package repository

import (
	"time"

	"github.com/bariasabda/kumparan/domain/entity"
)

func (r *repository) CreateArticle(article entity.Article) (*entity.Article, error) {
	article.Created = time.Now()
	err := r.db.Create(&article).Error

	if err != nil {
		return nil, err
	}

	return &article, nil
}
