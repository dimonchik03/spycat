package controllers

import (
	"net/http"
	"spycat/models"
	"spycat/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CatController struct {
	CatService *services.CatService
}

func NewCatController(catService *services.CatService) *CatController {
	return &CatController{CatService: catService}
}

func (ctrl *CatController) CreateCat(c *gin.Context) {
	var cat models.Cat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.CatService.CreateCat(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cat)
}

func (ctrl *CatController) GetCats(c *gin.Context) {
	cats, err := ctrl.CatService.GetCats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cats)
}

func (ctrl *CatController) GetCat(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	cat, err := ctrl.CatService.GetCat(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cat)
}

func (ctrl *CatController) UpdateCat(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var updatedCat models.Cat
	if err := c.ShouldBindJSON(&updatedCat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.CatService.UpdateCat(uint(id), &updatedCat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCat)
}

func (ctrl *CatController) DeleteCat(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := ctrl.CatService.DeleteCat(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cat deleted"})
}
