package services

import (
	"math"
	"todo-api/config"
	"todo-api/models"
)

// GetTasks fetches tasks with pagination and filter
func GetTasks(status string, search string, page int, limit int) ([]models.Task, int64, int, error) {
	var tasks []models.Task
	var totalTasks int64
	query := config.DB.Model(&models.Task{})

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&totalTasks).Error; err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalTasks) / float64(limit)))
	if totalPages == 0 {
		totalPages = 1
	}

	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, 0, err
	}

	return tasks, totalTasks, totalPages, nil
}
