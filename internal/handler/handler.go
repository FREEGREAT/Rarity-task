package handler

import (
	"github.com/gin-gonic/gin"

	"test_task/internal/handler/middleware"
	"test_task/internal/service"
)

type Handler struct {
	logger middleware.Logger
	api    service.RaribleApiClient
}

func NewHandler(log middleware.Logger, api service.RaribleApiClient) *Handler {
	return &Handler{
		logger: log,
		api:    api,
	}
}

func (h *Handler) InitRoutes(router *gin.Engine) *gin.Engine {
	routerGroup := router.Group("/nft")
	{
		routerGroup.GET("/ownerships/:id", h.getOwnershipById)
		routerGroup.POST("/trait-rarities", h.getTraitsRarity)
	}

	return router
}
