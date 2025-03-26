package userService

import "GorillaMuxProject/internal/web/users"

type UserService struct {
	repo UserRepository
}

func NewService(repo *UserStructRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user users.User) (users.User, error) {
	return s.repo.CreateUser(user)
}
func (s *UserService) GetAllUsers() ([]users.User, error) {
	return s.repo.GetAllUsers()
}
func (s *UserService) UpdateUserByID(id uint, user users.User) (users.User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
