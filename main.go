package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Spr1zze/barber-shop-backend/internal/db"
	migration "github.com/Spr1zze/barber-shop-backend/internal/db"
	Handler "github.com/Spr1zze/barber-shop-backend/internal/handlers"
	"github.com/Spr1zze/barber-shop-backend/internal/repository"
	"github.com/Spr1zze/barber-shop-backend/internal/services"
)

func main() {
	// 1. Connect to database
	sqlDB, err := db.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Run migrations
	err = migration.RunMigration(sqlDB)
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := db.ConnectGorm()
	if err != nil {
		log.Fatal(err)
	}

	// Dependency Injection typ sh in golang be like
	salonRepository := repository.NewSalonRepository(gormDB)
	salonService := services.NewSalonService(salonRepository)
	handler := Handler.NewHandler(salonService)

	// 3. Set up routes
	router := gin.Default()
	router.GET("/salons/:slug", handler.GetSalonBySlug) // slug is name example (downtown-hair)
	router.GET("/appointments/history", func(c *gin.Context) {
		Handler.GetAppointmentsHistory(c, sqlDB)
	})
	router.GET("/salon/details", func(c *gin.Context) {
		Handler.GetSalonDetails(c, sqlDB)
	})

	// 4. Start server LAST
	router.Run(":8000")
}
