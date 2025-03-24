package taskService

import (
	"GorillaMuxProject/internal/web/tasks"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task tasks.Task) (tasks.Task, error)
	GetAllTasks() ([]tasks.Task, error)
	UpdateTaskByID(id uint, task tasks.Task) (tasks.Task, error)
	DeleteTaskByID(id uint) error
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task tasks.Task) (tasks.Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return tasks.Task{}, result.Error
	}
	return task, nil
}
func (r *taskRepository) GetAllTasks() ([]tasks.Task, error) {
	var allTasks []tasks.Task
	err := r.db.Find(&allTasks).Error
	return allTasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task tasks.Task) (tasks.Task, error) {
	// Fetch the existing task from the database
	var existingTask tasks.Task
	err := r.db.First(&existingTask, id).Error
	if err != nil {
		return tasks.Task{}, err
	}

	// Update the fields of the existing task
	if task.Task != nil && *task.Task != "" {
		existingTask.Task = task.Task
	}
	// Update the IsDone field if it is provided
	if task.IsDone != nil { // Check if IsDone was provided
		existingTask.IsDone = task.IsDone
	}

	// Save the updated task back to the database
	err = r.db.Save(&existingTask).Error
	if err != nil {
		return tasks.Task{}, err
	}

	// Return the updated task
	return existingTask, nil
}
func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task tasks.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&task, id).Error
}
