// File: controllers/target_controller.go

package controllers

import (
	"net/http"
	"spycat/spycat/models"
	"spycat/spycat/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TargetController struct {
	targetService *services.TargetService
}

func NewTargetController(targetService *services.TargetService) *TargetController {
	return &TargetController{targetService: targetService}
}

func (ctrl *TargetController) AddTarget(c *gin.Context) {
	missionID, err := strconv.ParseUint(c.Param("mission_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
		return
	}

	var target models.Target
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	target.MissionID = uint(missionID)

	if err := ctrl.targetService.AddTarget(&target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, target)
}

func (ctrl *TargetController) UpdateTarget(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	target, err := ctrl.targetService.GetTargetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Target not found"})
		return
	}

	if err := c.ShouldBindJSON(target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.targetService.UpdateTarget(target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, target)
}

func (ctrl *TargetController) DeleteTarget(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := ctrl.targetService.DeleteTarget(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "target deleted"})
}

func (ctrl *TargetController) CompleteTarget(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := ctrl.targetService.CompleteTarget(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "target completed"})
}

func (ctrl *TargetController) UpdateNotes(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var noteUpdate struct {
		Notes string `json:"notes"`
	}
	if err := c.ShouldBindJSON(&noteUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.targetService.UpdateNotes(uint(id), noteUpdate.Notes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "notes updated"})
}

// File: controllers/target_controller.go

func (ctrl *TargetController) GetAllTargets(c *gin.Context) {
	targets, err := ctrl.targetService.GetAllTargets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, targets)
}

func (ctrl *TargetController) GetTarget(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid target ID"})
		return
	}

	target, err := ctrl.targetService.GetTargetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "target not found"})
		return
	}

	c.JSON(http.StatusOK, target)
}

func (ctrl *TargetController) GetTargetsByMission(c *gin.Context) {
	missionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
		return
	}

	targets, err := ctrl.targetService.GetTargetsByMissionID(uint(missionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, targets)
}
