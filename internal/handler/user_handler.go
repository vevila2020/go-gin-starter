package handler

import (
	"strconv"

	"go_basic_framework_template/internal/middleware"
	"go_basic_framework_template/internal/response"
	"go_basic_framework_template/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		service: service.NewUserService(),
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req service.CreateUserRequest
	if err := middleware.BindJSON(c, &req); err != nil {
		return
	}

	user, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	var req service.UpdateUserRequest
	if err := middleware.BindJSON(c, &req); err != nil {
		return
	}

	user, err := h.service.Update(c.Request.Context(), uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	user, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.NotFound(c, "user not found")
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) List(c *gin.Context) {
	var req service.UserListRequest
	if err := middleware.BindQuery(c, &req); err != nil {
		return
	}

	users, total, err := h.service.List(c.Request.Context(), &req)
	if err != nil {
		response.ServerError(c, "failed to get user list")
		return
	}

	response.SuccessPage(c, users, total, req.Page, req.PageSize)
}
