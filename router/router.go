package router

import (
	"github.com/elizabethrosales/blog/handler"
	"github.com/elizabethrosales/blog/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Router struct {
	*gin.Engine
}

func New(svc *service.Service, log *logrus.Logger) *Router {
	r := gin.New()
	r.ContextWithFallback = true

	// register routes
	registerRoutes(r, svc, log)

	return &Router{
		Engine: r,
	}
}

func registerRoutes(r *gin.Engine, svc *service.Service, log *logrus.Logger) {
	h := handler.New(svc, log)

	articlesRoutes := r.Group("articles")
	articlesRoutes.POST("", h.PostArticles)
	articlesRoutes.GET("", h.GetArticles)
	articlesRoutes.GET("/:id", h.GetArticle)
}
