package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	Query "github.com/Spr1zze/barber-shop-backend/internal/db"
)

func GetAppointmentsHistory(c *gin.Context, db *sql.DB) {
	history, err := Query.AppointmentHistory(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}
