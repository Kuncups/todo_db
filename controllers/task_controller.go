package controllers

import (
	"net/http"
	"strconv"
	"todo-api/config"
	"todo-api/models"
	"todo-api/services"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
)

// CreateTask untuk membuat task baru
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan task ke database
	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Task created successfully",
		"data":    task,
	})
}

func GetTasks(c *gin.Context) {
	status := c.DefaultQuery("status", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")

	tasks, totalTasks, totalPages, err := services.GetTasks(status, search, page, limit)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"pagination": gin.H{
			"current_page": page,
			"total_pages":  totalPages,
			"total_tasks":  totalTasks,
		},
	})
}

func GetTaskByID(c *gin.Context) {
	var task models.Task
	if err := config.DB.First(&task, c.Param("id")).Error; err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Task not found")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "Task retrieved successfully", task)
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	taskID := c.Param("id")

	// Cek apakah task dengan ID tersebut ada
	if err := config.DB.First(&task, taskID).Error; err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Task not found")
		return
	}

	// Bind JSON ke task (hanya mengupdate field yang diberikan)
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Update field sesuai request
	config.DB.Model(&task).Updates(updatedTask)

	// Response sesuai format yang diminta
	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    task,
	})
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := config.DB.First(&task, c.Param("id")).Error; err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Task not found")
		return
	}
	config.DB.Delete(&task)
	utils.RespondWithJSON(c, http.StatusOK, "Task deleted successfully", nil)
}
