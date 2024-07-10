package main

import (
	"log"
	"os"
	"spycat/spycat/config"
	"spycat/spycat/database"
	middleware2 "spycat/spycat/middleware"
	models2 "spycat/spycat/models"
	"spycat/spycat/routes"

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
	err = db.AutoMigrate(&models2.User{}, &models2.Cat{}, &models2.Mission{}, &models2.Target{})

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
	router.Use(middleware2.Logger())

	// integrate validation middleware
	router.Use(middleware2.Validator())
	router.Use(middleware2.AuthMiddleware())
	//register routes
	routes.RegisterRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
