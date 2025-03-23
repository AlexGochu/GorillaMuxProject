package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	// Fetch the existing task from the database
	var existingTask Task
	err := r.db.First(&existingTask, id).Error
	if err != nil {
		return Task{}, err
	}

	// Update the fields of the existing task
	if task.Task != "" {
		existingTask.Task = task.Task
	}
	existingTask.IsDone = task.IsDone

	// Save the updated task back to the database
	err = r.db.Save(&existingTask).Error
	if err != nil {
		return Task{}, err
	}

	// Return the updated task
	return existingTask, nil
}
func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&task, id).Error
}
