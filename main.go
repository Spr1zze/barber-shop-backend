package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Spr1zze/barber-shop-backend/internal/db"
	migration "github.com/Spr1zze/barber-shop-backend/internal/db"

	Handler "github.com/Spr1zze/barber-shop-backend/internal/handlers"
)

func main() {
	// 1. Connect to database
	db, err := db.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Run migrations
	err = migration.RunMigration(db)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Set up routes
	router := gin.Default()
	router.GET("/appointments/history", func(c *gin.Context) {
		Handler.GetAppointmentsHistory(c, db)
	})

	// 4. Start server LAST
	router.Run(":8000")
}
