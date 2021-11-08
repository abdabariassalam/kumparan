package repository

import (
	"github.com/bariasabda/kumparan/domain/entity"
)

func (r *repository) GetArticle(search string, author string) (*[]entity.Article, error) {
	var articles []entity.Article
	db := r.db.Where(`title like '%` + search + `%' or  body like '%` + search + `%'`)
	if author != "" {
		db.Where("author = ? ", author)
	}

	db.Find(&articles)

	err := db.Error
	if err != nil {
		return nil, err
	}

	return &articles, nil
}
