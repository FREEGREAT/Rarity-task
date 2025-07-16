package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getOwnershipById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}

	ownership, err := h.api.GetOwnershipByID(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("failed to get ownership: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ownership)
}
