package controller

import (
	"net/http"

	"github.com/siwanastudio/SA-65-SW/entity"
	"github.com/gin-gonic/gin"
)

// POST //watch_Tlts
func CreateProblemReport(c *gin.Context) {

	var ProblemReport entity.ProblemReport
	var Problem entity.Problem
	var Place_Class entity.Place_Class
	var Tlt entity.Toilet
	var RdZone entity.ReadingZone
	var ReschRoom entity.ResearchRoom
	var Com entity.Computer

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร ProblemReport
	if err := c.ShouldBindJSON(&ProblemReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Problem ด้วย id
	if tx := entity.DB().Where("id = ?", ProblemReport.Problem_ID).First(&Problem); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem not found"})
		return
	}

	//********************************************************************************
	// 10: ค้นด้วยไอดีชื่อสถานที่(PlaceName_id)
	if tx := entity.DB().Where("id = ?", ProblemReport.Tlt_id).First(&Tlt); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Toilet not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", ProblemReport.RdZone_id).First(&RdZone); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reding Zone not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", ProblemReport.ReschRoom_id).First(&ReschRoom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Research Room not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", ProblemReport.Com_id).First(&Com); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reding Zone not found"})
		return
	}
	//********************************************************************************

	// 11: ค้นหา Place_Class ด้วย id
	if tx := entity.DB().Where("id = ?", ProblemReport.Place_Class_ID).First(&Place_Class); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Place_Class not found"})
		return
	}

	// 12: สร้าง ProblemReport
	wv := entity.ProblemReport{
		Problem:     Problem,     // โยงความสัมพันธ์กับ Entity Problem
		Tlt:         Tlt,         // โยงความสัมพันธ์กับ Entity Toilet
		RdZone:      RdZone,      // โยงความสัมพันธ์กับ Entity ReadingZone
		ReschRoom:   ReschRoom,   // โยงความสัมพันธ์กับ Entity ResearchRoom
		Com:         Com,         // โยงความสัมพันธ์กับ Entity Computer
		Place_Class: Place_Class, // โยงความสัมพันธ์กับ Entity Place_Class
		/*
			WatchedTime: ProblemReport.WatchedTime, // ตั้งค่าฟิลด์ watchedTime
		*/
	}


	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

/*************************************************************************************************************************/

// GET /ProblemReport/:id *********************************************** !!!!
func GetProblemReport(c *gin.Context) {
	var ProblemReport entity.ProblemReport
	id := c.Param("id")
	if err := entity.DB().Preload("Problem").Preload("Place_Class").Preload("Toilet").Raw("SELECT * FROM watch_Toilet WHERE id = ?", id).Find(&ProblemReport).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ProblemReport})
}

// GET /watch_Tlts *********************************************** !!!!
func ListProblemReports(c *gin.Context) {
	var ProblemReports []entity.ProblemReport
	if err := entity.DB().Preload("Problem").Preload("Place_Class").Preload("Toilet").Raw("SELECT * FROM watch_Toilet").Find(&ProblemReports).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ProblemReports})
}

// DELETE /watch_Tlts/:id *********************************************** !!!!
func DeleteProblemReport(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM watch_Tlts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ProblemReport not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_Tlts *********************************************** !!!!
func UpdateProblemReport(c *gin.Context) {
	var ProblemReport entity.ProblemReport
	if err := c.ShouldBindJSON(&ProblemReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", ProblemReport.ID).First(&ProblemReport); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ProblemReport not found"})
		return
	}

	if err := entity.DB().Save(&ProblemReport).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ProblemReport})
}
