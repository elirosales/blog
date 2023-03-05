package service

import (
	"github.com/elizabethrosales/blog/database/models"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}

type Articles []Article

// GetArticle retrieves article by id
func (s *Service) GetArticle(id string) (response Articles, err error) {
	a := models.NewArticle(s.db)

	article, err := a.GetArticle(id)
	if err != nil {
		s.log.Errorf("Failed to retrieve article: %v", err.Error())
		return nil, err
	}

	response = append(response, fromArticleGORMToArticle(article))
	return response, err
}

// GetArticles retrieves all articles
func (s *Service) GetArticles() (response Articles, err error) {
	a := models.NewArticle(s.db)

	articles, err := a.GetArticles()
	if err != nil {
		s.log.Errorf("Failed to retrieve articles: %v", err.Error())
		return nil, err
	}

	for _, article := range articles {
		response = append(response, fromArticleGORMToArticle(article))
	}

	return response, err
}

func fromArticleGORMToArticle(data models.Article) Article {
	return Article{
		ID:      data.UUID,
		Title:   data.Title,
		Content: data.Content,
		Author:  data.Author,
	}
}
