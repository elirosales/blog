package handler

import (
	"net/http"
	"strings"

	"github.com/elizabethrosales/blog/utils"
	"github.com/gin-gonic/gin"
)

// GetArticle retrieves article by id
func (h *Handler) GetArticle(c *gin.Context) {
	id := c.Param("id")

	articles, err := h.svc.GetArticle(id)
	if err != nil {
		status := http.StatusBadRequest
		if strings.Contains(err.Error(), "not found") {
			status = http.StatusNotFound
		}

		utils.JSONErrResponse(c, status, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Success", articles)
	return
}

// GetArticles retrieves all articles
func (h *Handler) GetArticles(c *gin.Context) {
	articles, err := h.svc.GetArticles()
	if err != nil {
		utils.JSONErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Success", articles)
	return
}
