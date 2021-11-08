package service

func (s *service) GetArticle(search string, author string) (*[]Article, error) {
	var resp []Article
	result, err := s.repo.GetArticle(search, author)

	for _, article := range *result {
		resp = append(resp, Article{
			Id:      article.Id,
			Author:  article.Author,
			Title:   article.Title,
			Body:    article.Body,
			Created: article.Created,
		})
	}

	return &resp, err
}
