package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/amliyanage/go-jwt-tasks/models"
	"github.com/amliyanage/go-jwt-tasks/repo"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")

		var req CreateTaskRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task := models.Task{
			Title:       req.Title,
			Description: req.Description,
			UserID:      userID,
		}

		if err := repo.DB.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
			return
		}

		c.JSON(http.StatusCreated, task)
	}
}

func GetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")

		var tasks []models.Task
		if err := repo.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
			return
		}

		c.JSON(http.StatusOK, tasks)
	}
}

func GetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
			return
		}

		var task models.Task
		if err := repo.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

func UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
			return
		}

		var task models.Task
		if err := repo.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}

		var req UpdateTaskRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Title != nil {
			task.Title = *req.Title
		}
		if req.Description != nil {
			task.Description = *req.Description
		}
		if req.Completed != nil {
			task.Completed = *req.Completed
		}

		if err := repo.DB.Save(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
			return
		}

		result := repo.DB.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
	}
}
