package taskService

import "GorillaMuxProject/internal/web/tasks"

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task tasks.Task) (tasks.Task, error) {
	return s.repo.CreateTask(task)
}
func (s *TaskService) GetAllTasks() ([]tasks.Task, error) {
	return s.repo.GetAllTasks()
}
func (s *TaskService) UpdateTaskByID(id uint, task tasks.Task) (tasks.Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
