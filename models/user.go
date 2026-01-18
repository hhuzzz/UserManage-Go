package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(20)"`
	Age       int       `json:"age" gorm:"type:int"`
	Status    int       `json:"status" gorm:"type:int;default:1"` // 1: active, 0: inactive
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required,email"`
	Phone  string `json:"phone"`
	Age    int    `json:"age"`
	Status int    `json:"status"`
}

type UpdateUserRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email" binding:"omitempty,email"`
	Phone  string `json:"phone"`
	Age    int    `json:"age"`
	Status int    `json:"status"`
}
