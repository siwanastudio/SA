package controller

import (
	"net/http"

	"github.com/siwanastudio/SA-65-SW/entity"
	"github.com/gin-gonic/gin"
)

// POST /Toilets
func CreateToilet(c *gin.Context) {
	var Toilet entity.Toilet
	if err := c.ShouldBindJSON(&Toilet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Toilet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Toilet})
}

// GET /Toilet/:id
func GetToilet(c *gin.Context) {
	var Toilet entity.Toilet

	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM Toilets WHERE id = ?", id).Find(&Toilet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Toilet})
}

// GET /Toilets
func ListToilets(c *gin.Context) {
	var Toilets []entity.Toilet
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM Toilets").Find(&Toilets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Toilets})
}

// DELETE /Toilets/:id
func DeleteToilet(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Toilets WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Toilet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Toilets
func UpdateToilet(c *gin.Context) {
	var Toilet entity.Toilet
	if err := c.ShouldBindJSON(&Toilet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Toilet.ID).First(&Toilet); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Toilet not found"})
		return
	}

	if err := entity.DB().Save(&Toilet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Toilet})
}
