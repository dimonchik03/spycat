package controllers

import (
	"fmt"
	"net/http"
	"spycat/spycat/models"
	services2 "spycat/spycat/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MissionController struct {
	MissionService *services2.MissionService
	CatService     *services2.CatService
}

func NewMissionController(missionService *services2.MissionService, catService *services2.CatService) *MissionController {
	return &MissionController{MissionService: missionService, CatService: catService}
}

func (ctrl *MissionController) CreateMission(c *gin.Context) {
	var mission models.Mission
	if err := c.ShouldBindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.CatService.GetCat(mission.CatID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CatID"})
		return
	}

	if err := ctrl.MissionService.CreateMission(&mission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mission)
}

func (ctrl *MissionController) GetMissions(c *gin.Context) {
	missions, err := ctrl.MissionService.GetMissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, missions)
}

func (ctrl *MissionController) GetMission(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	mission, err := ctrl.MissionService.GetMission(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mission)
}

func (ctrl *MissionController) UpdateMission(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	existingMission, err := ctrl.MissionService.GetMission(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
		return
	}

	if err := c.ShouldBindJSON(existingMission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("updated mission data: %+v\n", existingMission)

	// Validate CatID only if it's provided in the request
	if existingMission.CatID != 0 {
		if _, err := ctrl.CatService.GetCat(existingMission.CatID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid CatID: %v", err)})
			return
		}
	}

	if err := ctrl.MissionService.UpdateMission(uint(id), existingMission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingMission)
}

func (ctrl *MissionController) DeleteMission(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := ctrl.MissionService.DeleteMission(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "mission deleted"})
}

func (ctrl *MissionController) CompleteMission(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := ctrl.MissionService.CompleteMission(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "mission completed"})
}
