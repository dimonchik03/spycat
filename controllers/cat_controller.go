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
	cat, exists := c.Get("cat")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cat not found in context"})
		return
	}

	if err := ctrl.CatService.CreateCat(cat.(*models.Cat)); err != nil {
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
	updatedCat, exists := c.Get("cat")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cat not found in context"})
		return
	}

	if err := ctrl.CatService.UpdateCat(uint(id), updatedCat.(*models.Cat)); err != nil {
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
