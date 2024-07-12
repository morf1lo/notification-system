package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morf1lo/notification-system/internal/feed/model"
)

func (h *Handler) feedPublish(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}

	if err := h.services.Feed.Publish(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "error": nil})
}
