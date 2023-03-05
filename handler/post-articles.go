package handler

import (
	"net/http"

	"github.com/elizabethrosales/blog/service"
	"github.com/elizabethrosales/blog/utils"
	"github.com/gin-gonic/gin"
)

// PostArticles creates an article
func (h *Handler) PostArticles(c *gin.Context) {
	var req service.PostArticlesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Invalid request: %v", err.Error())

		utils.JSONErrResponse(c, http.StatusBadRequest, "Invalid request", h.translateBindingErr(err))
		return
	}

	article, err := h.svc.PostArticle(req)
	if err != nil {
		utils.JSONErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Success", article)
	return
}
