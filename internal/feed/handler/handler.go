package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/morf1lo/notification-system/internal/feed/service"
	"github.com/spf13/viper"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{viper.GetString("web.origin")},
		AllowMethods: []string{"POST"},
	}))

	api := router.Group("/api")
	{
		feed := api.Group("/feed")
		{
			feed.POST("/publish", h.mwAdmin, h.feedPublish)
		}
	}

	return router
}
