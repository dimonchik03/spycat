package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"spycat/controllers"
	"spycat/repositories"
	"spycat/services"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	catRepo := repositories.NewCatRepository(db)
	catService := services.NewCatService(catRepo)
	catController := controllers.NewCatController(catService)

	missionRepo := repositories.NewMissionRepository(db)
	missionService := services.NewMissionService(missionRepo)
	missionController := controllers.NewMissionController(missionService, catService)

	targetRepo := repositories.NewTargetRepository(db)
	targetService := services.NewTargetService(targetRepo, missionRepo)
	targetController := controllers.NewTargetController(targetService)

	authController := controllers.NewAuthController(db)

	// сat routes
	router.POST("/cats", catController.CreateCat)
	router.GET("/cats", catController.GetCats)
	router.GET("/cats/:id", catController.GetCat)
	router.PUT("/cats/:id", catController.UpdateCat)
	router.DELETE("/cats/:id", catController.DeleteCat)

	// mission routes
	router.POST("/missions", missionController.CreateMission)
	router.GET("/missions", missionController.GetMissions)
	router.GET("/missions/:id", missionController.GetMission)
	router.PUT("/missions/:id", missionController.UpdateMission)
	router.DELETE("/missions/:id", missionController.DeleteMission)
	router.PUT("/missions/:id/complete", missionController.CompleteMission)

	// target routes
	router.POST("/missions/:id/targets", targetController.AddTarget)
	router.PUT("/targets/:id", targetController.UpdateTarget)
	router.DELETE("/targets/:id", targetController.DeleteTarget)
	router.PUT("/targets/:id/complete", targetController.CompleteTarget)
	router.PUT("/targets/:id/notes", targetController.UpdateNotes)
	router.GET("/targets", targetController.GetAllTargets)
	router.GET("/targets/:id", targetController.GetTarget)
	router.GET("/missions/:id/targets", targetController.GetTargetsByMission)

	router.POST("/login", authController.Login)

}
