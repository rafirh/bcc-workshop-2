package repository

import (
	"gorm.io/gorm"
	"context"
	"bcc-workshop-2/internal/entity"
)

type IRestaurantRepository interface {
	CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) error
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

