package service

import (
	"context"
	"errors"

	"go_basic_framework_template/internal/model"
	"go_basic_framework_template/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	Nickname string `json:"nickname"`
}

type UpdateUserRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Status   *int   `json:"status"`
}

type UserListRequest struct {
	Page     int `form:"page" validate:"min=1"`
	PageSize int `form:"page_size" validate:"min=1,max=100"`
}

func (s *UserService) Create(ctx context.Context, req *CreateUserRequest) (*model.User, error) {
	existingUser, _ := s.repo.GetByUsername(ctx, req.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	existingEmail, _ := s.repo.GetByEmail(ctx, req.Email)
	if existingEmail != nil {
		return nil, errors.New("email already exists")
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
		Status:   1,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, id uint, req *UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Status != nil {
		user.Status = *req.Status
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) List(ctx context.Context, req *UserListRequest) ([]model.User, int64, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	return s.repo.List(ctx, req.Page, req.PageSize)
}
