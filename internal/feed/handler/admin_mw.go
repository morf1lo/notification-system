package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) mwAdmin(c *gin.Context) {
	accessToken := strings.TrimSpace(c.GetHeader("Access"))

	if accessToken != os.Getenv("ACCESS_TOKEN") {
		c.JSON(http.StatusForbidden, gin.H{"ok": false, "error": "you have no access"})
		c.Abort()
		return
	}

	c.Next()
}
