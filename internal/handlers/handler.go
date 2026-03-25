package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Spr1zze/barber-shop-backend/internal/repository"
	"github.com/Spr1zze/barber-shop-backend/internal/services"
)

type Handler struct {
	salonService *services.SalonService
}

func NewHandler(salonService *services.SalonService) *Handler {
	return &Handler{salonService: salonService}
}

func (h *Handler) GetSalonBySlug(c *gin.Context) {
	salon, err := h.salonService.GetSalonPage(c.Request.Context(), c.Param("slug"))
	if err != nil {
		if errors.Is(err, repository.ErrSalonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "salon not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load salon"})
		return
	}

	c.JSON(http.StatusOK, salon)
}
