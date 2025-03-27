package userService

import (
	"GorillaMuxProject/internal/web/users"
	"errors"
	"fmt"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"
	"strings"
)

type UserRepository interface {
	CreateUser(user users.User) (users.User, error)
	GetAllUsers() ([]users.User, error)
	UpdateUserByID(id uint, user users.User) (users.User, error)
	DeleteUserByID(id uint) error
	GetTasksByUserID(id uint) ([]users.Task, error)
}
type UserStructRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserStructRepository {
	return &UserStructRepository{db: db}
}

func validateEmail(email openapi_types.Email) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	// Basic email format check
	if !strings.Contains(string(email), "@") {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func (r *UserStructRepository) CreateUser(user users.User) (users.User, error) {
	if user.Email != nil && *user.Email != "" {
		if err := validateEmail(*user.Email); err != nil {
			return users.User{}, fmt.Errorf("invalid email: %w", err)
		}
	}
	userToCreate := users.User{
		Email:    user.Email,
		Password: user.Password, // Hash this!
	}

	if err := r.db.Create(&userToCreate).Error; err != nil {
		return users.User{}, err
	}
	return users.User{
		Id:    userToCreate.Id,
		Email: userToCreate.Email,
	}, nil
}
func (r *UserStructRepository) GetAllUsers() ([]users.User, error) {
	var allUsers []users.User
	err := r.db.Find(&allUsers).Error
	return allUsers, err
}

func (r *UserStructRepository) UpdateUserByID(id uint, user users.User) (users.User, error) {
	// Fetch the existing user from the database
	var existingUser users.User
	err := r.db.First(&existingUser, id).Error
	if err != nil {
		return users.User{}, err
	}

	// Update only the allowed fields
	if user.Email != nil && *user.Email != "" {
		if err := validateEmail(*user.Email); err != nil {
			return users.User{}, fmt.Errorf("invalid email: %w", err)
		}
		existingUser.Email = user.Email
	}

	// Save the updated user back to the database
	err = r.db.Save(&existingUser).Error
	if err != nil {
		return users.User{}, err
	}

	// Return the updated user
	return users.User{
		Id:    existingUser.Id,
		Email: existingUser.Email,
	}, nil
}
func (r *UserStructRepository) DeleteUserByID(id uint) error {
	var user users.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&user, id).Error
}
func (r *UserStructRepository) GetTasksByUserID(id uint) ([]users.Task, error) {
	var tasks []users.Task

	// First check if user exists
	var user users.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	// Using explicit join query
	err := r.db.Where("user_id = ?", id).Find(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %v", err)
	}

	return tasks, nil
}
