package main

import (
	"log"
	"os"
	"spycat/config"
	"spycat/database"
	"spycat/middleware"
	"spycat/models"
	"spycat/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load environment variables
	config.LoadEnv()

	// initialize the database connection
	dsn := config.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	//  database migration
	err = db.AutoMigrate(&models.User{}, &models.Cat{}, &models.Mission{}, &models.Target{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	err = database.SeedDatabase(db)
	if err != nil {
		log.Fatalf("failed to seed database: %v", err)
	}

	//initialize router
	router := gin.Default()

	//integrate logging middleware
	router.Use(middleware.Logger())

	// integrate validation middleware
	router.Use(middleware.Validator())
	router.Use(middleware.AuthMiddleware())
	//register routes
	routes.RegisterRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
