package service

import (
	"github.com/elizabethrosales/blog/database/models"
	"github.com/google/uuid"
)

type PostArticlesRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author  string `json:"author" binding:"required"`
}

// PostArticles creates an article
func (s *Service) PostArticle(request PostArticlesRequest) (response Article, err error) {
	a := models.NewArticle(s.db)

	id := uuid.NewString()
	if err = a.InsertArticle(models.Article{
		UUID:    id,
		Title:   request.Title,
		Content: request.Content,
		Author:  request.Author,
	}); err != nil {
		s.log.Errorf("Failed to insert article: %v", err.Error())
		return response, err
	}

	response.ID = id
	return response, err
}
