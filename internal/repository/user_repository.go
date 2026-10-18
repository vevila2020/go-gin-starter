package repository

import (
	"context"
	"errors"

	"go-gin-starter/internal/database"
	"go-gin-starter/internal/model"

	"gorm.io/gorm"
)

var ErrDatabaseDisabled = errors.New("database is disabled")

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) db(ctx context.Context) (*gorm.DB, error) {
	if !database.IsEnabled() {
		return nil, ErrDatabaseDisabled
	}
	return database.DB.WithContext(ctx), nil
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	db, err := r.db(ctx)
	if err != nil {
		return err
	}
	return db.Create(user).Error
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	db, err := r.db(ctx)
	if err != nil {
		return err
	}
	return db.Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	db, err := r.db(ctx)
	if err != nil {
		return err
	}
	return db.Delete(&model.User{}, id).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	db, err := r.db(ctx)
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	db, err := r.db(ctx)
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	db, err := r.db(ctx)
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) List(ctx context.Context, page, pageSize int) ([]model.User, int64, error) {
	db, err := r.db(ctx)
	if err != nil {
		return nil, 0, err
	}

	var users []model.User
	var total int64

	query := db.Model(&model.User{})
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}