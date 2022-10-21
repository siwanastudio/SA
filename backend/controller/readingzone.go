package controller

import (
	"net/http"

	"github.com/siwanastudio/SA-65-SW/entity"
	"github.com/gin-gonic/gin"
)

// POST /ReadingZones
func CreateReadingZone(c *gin.Context) {
	var ReadingZone entity.ReadingZone
	if err := c.ShouldBindJSON(&ReadingZone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&ReadingZone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ReadingZone})
}

// GET /ReadingZone/:id
func GetReadingZone(c *gin.Context) {
	var ReadingZone entity.ReadingZone

	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM ReadingZones WHERE id = ?", id).Find(&ReadingZone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ReadingZone})
}

// GET /ReadingZones
func ListReadingZones(c *gin.Context) {
	var ReadingZones []entity.ReadingZone
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM ReadingZones").Find(&ReadingZones).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ReadingZones})
}

// DELETE /ReadingZones/:id
func DeleteReadingZone(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ReadingZones WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReadingZone not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ReadingZones
func UpdateReadingZone(c *gin.Context) {
	var ReadingZone entity.ReadingZone
	if err := c.ShouldBindJSON(&ReadingZone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", ReadingZone.ID).First(&ReadingZone); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReadingZone not found"})
		return
	}

	if err := entity.DB().Save(&ReadingZone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ReadingZone})
}
