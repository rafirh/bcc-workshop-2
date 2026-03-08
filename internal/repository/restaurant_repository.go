package repository

import (
	"bcc-workshop-2/internal/entity"
	"bcc-workshop-2/internal/model"
	"context"
	"gorm.io/gorm"
)

type IRestaurantRepository interface {
	CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) error
	GetAllRestaurants(ctx context.Context, pagination model.Pagination) ([]entity.Restaurant, error)
}

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) IRestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) error {
	if err := r.db.WithContext(ctx).Create(&restaurant).Error; err != nil {
		return err
	}
	return nil
}

func (r *RestaurantRepository) GetAllRestaurants(ctx context.Context, pagination model.Pagination) ([]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
	offset := (pagination.Page - 1) * pagination.Limit
	if err := r.db.WithContext(ctx).Order("created_at DESC").Limit(pagination.Limit).Offset(offset).Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}
