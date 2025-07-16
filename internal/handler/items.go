package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"test_task/internal/dto"
)

func (h *Handler) getTraitsRarity(c *gin.Context) {
	var req dto.TraitRarityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
		return
	}

	resp, err := h.api.GetNftTraitsRarity(c.Request.Context(), req.CollectionId, req.Properties)
	if err != nil {
		h.logger.Error("Failed to get traits rarity: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
