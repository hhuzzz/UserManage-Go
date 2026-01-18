package services

import (
	"errors"
	"hello/models"
	"hello/repositories"
)

type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// Check if email already exists
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	user := &models.User{
		Name:   req.Name,
		Email:  req.Email,
		Phone:  req.Phone,
		Age:    req.Age,
		Status: req.Status,
	}

	if req.Status == 0 {
		user.Status = 1 // Default to active
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Check if email is being changed and if it conflicts with another user
	if req.Email != "" && req.Email != user.Email {
		existingUser, err := s.repo.FindByEmail(req.Email)
		if err == nil && existingUser != nil && existingUser.ID != id {
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Age != 0 {
		user.Age = req.Age
	}
	if req.Status != 0 {
		user.Status = req.Status
	}

	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
