package services

import (
	"errors"
	"hello/models"
	"hello/repositories"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id uint) error
	SearchUsers(name string, page, size int, sortBy, sortOrder string) ([]models.User, int64, error)
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

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Phone:    req.Phone,
		Age:      req.Age,
		Status:   req.Status,
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

func (s *userService) SearchUsers(name string, page, size int, sortBy, sortOrder string) ([]models.User, int64, error) {
	name = strings.TrimSpace(name)
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	sortBy = strings.ToLower(strings.TrimSpace(sortBy))
	allowedSort := map[string]struct{}{
		"id":         {},
		"name":       {},
		"email":      {},
		"phone":      {},
		"age":        {},
		"status":     {},
		"created_at": {},
	}
	if _, ok := allowedSort[sortBy]; !ok {
		sortBy = "created_at"
	}

	sortOrder = strings.ToLower(strings.TrimSpace(sortOrder))
	sortDesc := true
	if sortOrder == "asc" {
		sortDesc = false
	}

	return s.repo.SearchByName(name, page, size, sortBy, sortDesc)
}
