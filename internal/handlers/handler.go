package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Spr1zze/barber-shop-backend/internal/repository"
	"github.com/Spr1zze/barber-shop-backend/internal/services"
	Type "github.com/Spr1zze/barber-shop-backend/model"
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

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func PostUsers(c *gin.Context) {
	var newUser Type.User

	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

var users = []Type.User{
	{ID: "1", Name: "Camilla Frederiksen", Phone: 23406751},
	{ID: "2", Name: "Johan Henriksen", Phone: 57279182},
	{ID: "3", Name: "Peter Knudsen", Phone: 12749501},
	{ID: "4", Name: "Lone Benedict", Phone: 98347261},
}
