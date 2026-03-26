package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Spr1zze/barber-shop-backend/internal/repository"
	"github.com/Spr1zze/barber-shop-backend/internal/services"
	Type "github.com/Spr1zze/barber-shop-backend/model"
)

type Handler struct {
	salonService   *services.SalonService
	bookingService *services.BookingService
}

func NewHandler(salonService *services.SalonService, bookingService *services.BookingService) *Handler {
	return &Handler{
		salonService:   salonService,
		bookingService: bookingService,
	}
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

func (h *Handler) ListSalonBarbers(c *gin.Context) {
	barbers, err := h.bookingService.ListBarbers(c.Request.Context(), c.Param("slug"))
	if err != nil {
		respondWithBookingError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"barbers": barbers})
}

func (h *Handler) GetBarberAvailability(c *gin.Context) {
	barberID := c.Query("barberId")
	serviceID := c.Query("serviceId")
	dateValue := c.Query("date")
	if barberID == "" || serviceID == "" || dateValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "barberId, serviceId and date are required"})
		return
	}

	day, err := time.Parse("2006-01-02", dateValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date must be formatted as YYYY-MM-DD"})
		return
	}

	slots, err := h.bookingService.GetAvailability(c.Request.Context(), c.Param("slug"), barberID, serviceID, day)
	if err != nil {
		respondWithBookingError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"slots": slots})
}

func (h *Handler) CreateBooking(c *gin.Context) {
	var payload Type.BookingRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid booking fields"})
		return
	}

	confirmation, err := h.bookingService.CreateBooking(c.Request.Context(), c.Param("slug"), &payload)
	if err != nil {
		respondWithBookingError(c, err)
		return
	}

	c.JSON(http.StatusCreated, confirmation)
}

func respondWithBookingError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, repository.ErrSalonNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "salon not found"})
	case errors.Is(err, repository.ErrBarberNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "barber not found"})
	case errors.Is(err, repository.ErrServiceNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
	case errors.Is(err, services.ErrBarberNotAssigned):
		c.JSON(http.StatusBadRequest, gin.H{"error": "barber does not work at this salon"})
	case errors.Is(err, services.ErrServiceNotInSalon):
		c.JSON(http.StatusBadRequest, gin.H{"error": "service is not offered at this salon"})
	case errors.Is(err, services.ErrSalonClosedOnDay):
		c.JSON(http.StatusBadRequest, gin.H{"error": "salon is closed that day"})
	case errors.Is(err, services.ErrOutsideOpeningHours):
		c.JSON(http.StatusBadRequest, gin.H{"error": "time is outside opening hours"})
	case errors.Is(err, services.ErrSlotUnavailable):
		c.JSON(http.StatusConflict, gin.H{"error": "time slot already taken"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "booking request failed"})
	}
}
