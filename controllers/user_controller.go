package controllers

import (
	"hello/models"
	"hello/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) IndexPage(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": "Failed to load users",
		})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"users": users,
	})
}

func (c *UserController) CreatePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form.html", gin.H{
		"title": "Create User",
		"action": "/users",
		"method": "POST",
		"user":   nil,
	})
}

func (c *UserController) EditPage(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	user, err := c.service.GetUserByID(uint(id))
	if err != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	ctx.HTML(http.StatusOK, "form.html", gin.H{
		"title": "Edit User",
		"action": "/users/" + ctx.Param("id"),
		"method": "PUT",
		"user":   user,
	})
}

func (c *UserController) DetailPage(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	user, err := c.service.GetUserByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusNotFound, "user_detail.html", gin.H{
			"error": "User not found",
		})
		return
	}

	ctx.HTML(http.StatusOK, "user_detail.html", gin.H{
		"user": user,
	})
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.CreateUser(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) SearchUsers(ctx *gin.Context) {
	name := ctx.Query("name")
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size"})
		return
	}
	sortBy := ctx.DefaultQuery("sort_by", "created_at")
	sortOrder := ctx.DefaultQuery("sort_order", "desc")

	users, total, err := c.service.SearchUsers(name, page, size, sortBy, sortOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"items": users,
		"page":  page,
		"size":  size,
		"total": total,
	})
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.service.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.UpdateUser(uint(id), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.service.DeleteUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
