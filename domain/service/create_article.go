package service

import (
	"errors"
	"strings"

	"github.com/bariasabda/kumparan/domain/entity"
)

func (s *service) CreateArticle(article entity.Article) (*Article, error) {
	resp, err := s.repo.CreateArticle(article)
	if err != nil {
        if strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
            return nil, errors.New("Author and Title already in database")
        }
		return nil, err
	}

	return &Article{
		Id:      resp.Id,
		Author:  resp.Author,
		Title:   resp.Title,
		Body:    resp.Body,
		Created: resp.Created,
	}, err
}
