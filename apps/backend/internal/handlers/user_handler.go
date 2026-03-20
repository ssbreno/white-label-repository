package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssbreno/white-label-repository/backend/internal/database"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
	"github.com/ssbreno/white-label-repository/backend/internal/repositories"
	"github.com/ssbreno/white-label-repository/backend/internal/services"
)

type userHandler struct {
	service *services.UserService
}

func newUserHandler(db *database.DB) *userHandler {
	repo := repositories.NewUserRepository(db)
	svc := services.NewUserService(repo)
	return &userHandler{service: svc}
}

// list handles GET /api/users
func (h *userHandler) list(c *gin.Context) {
	users, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError[any](err.Error()))
		return
	}
	if users == nil {
		users = []models.User{}
	}
	c.JSON(http.StatusOK, models.NewSuccess(users, "Users retrieved"))
}

// getByID handles GET /api/users/:id
func (h *userHandler) getByID(c *gin.Context) {
	user, err := h.service.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewError[any](err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewSuccess(user, "User retrieved"))
}

// create handles POST /api/users
func (h *userHandler) create(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewError[any](err.Error()))
		return
	}
	user, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError[any](err.Error()))
		return
	}
	c.JSON(http.StatusCreated, models.NewSuccess(user, "User created"))
}

// update handles PUT /api/users/:id
func (h *userHandler) update(c *gin.Context) {
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewError[any](err.Error()))
		return
	}
	user, err := h.service.Update(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError[any](err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewSuccess(user, "User updated"))
}

// delete handles DELETE /api/users/:id
func (h *userHandler) delete(c *gin.Context) {
	if err := h.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, models.NewError[any](err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewSuccess[any](nil, "User deleted"))
}
