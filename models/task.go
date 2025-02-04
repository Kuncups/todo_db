package models

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"oneof=pending completed"`
	DueDate     string `json:"due_date" binding:"required" gorm:"type:date"`
	// CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	// UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
