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
	router.GET("/users", Handler.GetUsers)
	router.POST("/users", Handler.PostUsers)
	router.GET("/salons/:slug", handler.GetSalonBySlug) // slug is name example (downtown-hair)

	// 4. Start server LAST
	router.Run(":8000")
}
