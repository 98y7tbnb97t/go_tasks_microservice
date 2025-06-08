package task

import (
	db "github.com/98y7tbnb97t/tasks-service/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetTasksFromBd() ([]Task, error) {
	var tasks []Task
	result := db.DB.Omit("created_at", "updated_at", "deleted_at").Find(&tasks)
	return tasks, result.Error
}

func (r *Repository) CreateTaskFromBd(task *Task) error {
	return db.DB.Create(task).Error
}

func (r *Repository) UpdateTaskFromBd(id string, task *Task) error {
	var existingTask Task
	if err := db.DB.First(&existingTask, id).Error; err != nil {
		return err
	}
	return db.DB.Save(task).Error
}

func (r *Repository) PatchTaskFromBd(id string, task *Task) error {
	var existingTask Task
	if err := db.DB.First(&existingTask, id).Error; err != nil {
		return err
	}
	return db.DB.Save(task).Error
}

func (r *Repository) DeleteTaskFromBd(id string) error {
	return db.DB.Delete(&Task{}, id).Error
}

func (r *Repository) GetTaskByIDFromBd(id string, task *Task) error {
	return db.DB.Omit("created_at", "updated_at", "deleted_at").First(task, "id = ? AND deleted_at IS NULL", id).Error
}

// GetTasksForUser retrieves all tasks belonging to a specific user
func (r *Repository) GetTasksForUserFromBd(userID uint) ([]Task, error) {
	var tasks []Task
	// Omit the _at fields from the result
	result := db.DB.Omit("created_at", "updated_at", "deleted_at").Where("user_id = ?", userID).Find(&tasks)
	return tasks, result.Error
}
