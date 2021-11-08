package service

import (
	"time"

	"github.com/bariasabda/kumparan/config"
	"github.com/bariasabda/kumparan/domain/entity"
	"github.com/bariasabda/kumparan/domain/repository"
)

type service struct {
	cfg  config.Config
	repo repository.RepositoryInterface
}

type ServiceInterface interface {
	GetArticle(search string, author string) (*[]Article, error)
	CreateArticle(article entity.Article) (*Article, error)
}

type Article struct {
	Id      int       `json:"id"`
	Author  string    `json:"author"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

func NewService(config config.Config, repo repository.RepositoryInterface) *service {
	return &service{
		cfg:  config,
		repo: repo,
	}
}
