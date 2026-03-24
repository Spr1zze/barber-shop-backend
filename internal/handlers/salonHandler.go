package handlers

import(
	"net/http"

	"github.com/gin-gonic/gin"
	"database/sql"
	Query "github.com/Spr1zze/barber-shop-backend/internal/db"
)

func GetSalonDetails(c *gin.Context,db *sql.DB)  {
	
	salon, err := Query.SalonDetails(db)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} 

	c.JSON(http.StatusOK, salon)
}
