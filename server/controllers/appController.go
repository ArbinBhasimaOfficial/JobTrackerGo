package controllers

import (
	"net/http"
	"time"

	"github.com/ArbinBhasimaOfficial/JobTracker/models"

	"github.com/ArbinBhasimaOfficial/JobTracker/config"

	"github.com/gin-gonic/gin"
)

type ApplicationInput struct {
	CompanyName string `json:"company_name" binding:"required,min=2"`
	JobTitle    string `json:"job_title" binding:"required"`
	JobType     string `json:"job_type" binding:"required,oneof=Internship Full-time Part-time"`
	Status      string `json:"status" binding:"required,oneof=Applied Interviewing Offer Rejected"`
	AppliedDate string `json:"applied_date" binding:"required" time_format:"2006-01-02"`
	Notes       string `json:"notes"`
}

// GET /applications
func GetApplications(c *gin.Context) {
	var apps []models.Application
	query := config.DB.Model(&models.Application{})

	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if search := c.Query("search"); search != "" {
		query = query.Where("company_name ILIKE ? OR job_title ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Order by most recent applied date
	if err := query.Order("applied_date desc").Find(&apps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records"})
		return
	}

	c.JSON(http.StatusOK, apps)
}

// GET /applications/:id
func GetApplicationByID(c *gin.Context) {
	var app models.Application
	id := c.Param("id")

	if err := config.DB.First(&app, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	c.JSON(http.StatusOK, app)
}

// POST /applications
func CreateApplication(c *gin.Context) {
	var input ApplicationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", input.AppliedDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	app := models.Application{
		CompanyName: input.CompanyName,
		JobTitle:    input.JobTitle,
		JobType:     models.JobType(input.JobType),
		Status:      models.Status(input.Status),
		AppliedDate: parsedDate,
		Notes:       input.Notes,
	}

	if err := config.DB.Create(&app).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save application"})
		return
	}

	c.JSON(http.StatusCreated, app)
}

// PATCH /applications/:id
func UpdateApplication(c *gin.Context) {
	id := c.Param("id")
	var app models.Application

	if err := config.DB.First(&app, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	var input ApplicationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", input.AppliedDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	updates := models.Application{
		CompanyName: input.CompanyName,
		JobTitle:    input.JobTitle,
		JobType:     models.JobType(input.JobType),
		Status:      models.Status(input.Status),
		AppliedDate: parsedDate,
		Notes:       input.Notes,
	}

	config.DB.Model(&app).Updates(updates)
	c.JSON(http.StatusOK, app)
}

// DELETE /applications/:id
func DeleteApplication(c *gin.Context) {
	id := c.Param("id")
	var app models.Application

	if err := config.DB.First(&app, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	config.DB.Delete(&app)
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}
