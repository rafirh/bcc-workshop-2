package repository

import "gorm.io/gorm"

type Repository struct {
	RestaurantRepository IRestaurantRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		RestaurantRepository: NewRestaurantRepository(db),
	}
}
