package controller

import (
	"net/http"
	"github.com/siwanastudio/SA-65-SW/entity"
	"github.com/gin-gonic/gin"
)

// POST /Place_Classs
func CreatePlace_Class(c *gin.Context) {
	var Place_Class entity.Place_Class
	if err := c.ShouldBindJSON(&Place_Class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Place_Class).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Place_Class})
}

// GET /Place_Class/:id
func GetPlace_Class(c *gin.Context) {
	var Place_Class entity.Place_Class
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Place_Class WHERE id = ?", id).Scan(&Place_Class).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Place_Class})
}

// GET /Place_Classs
func ListPlace_Class(c *gin.Context) {
	var Place_Classs []entity.Place_Class
	if err := entity.DB().Raw("SELECT * FROM Place_Classs").Scan(&Place_Classs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Place_Classs})
}

// DELETE /Place_Classs/:id
func DeletePlace_Class(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Place_Classs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Place_Class not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Place_Classs
func UpdatePlace_Class(c *gin.Context) {
	var Place_Class entity.Place_Class
	if err := c.ShouldBindJSON(&Place_Class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Place_Class.ID).First(&Place_Class); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Place_Class not found"})
		return
	}

	if err := entity.DB().Save(&Place_Class).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Place_Class})
}
