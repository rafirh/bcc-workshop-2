package usecase

import (
	"context"

	"bcc-workshop-2/internal/entity"
	"bcc-workshop-2/internal/model"
	"bcc-workshop-2/internal/repository"

	"github.com/google/uuid"
)

type IRestaurantUsecase interface {
	CreateRestaurant(ctx context.Context, req model.CreateRestaurant) error
}

type RestaurantUseCase struct {
	repo repository.IRestaurantRepository
}

func NewRestaurantUseCase(repo repository.IRestaurantRepository) IRestaurantUsecase {
	return &RestaurantUseCase{repo: repo}
}

func (uc *RestaurantUseCase) CreateRestaurant(ctx context.Context, req model.CreateRestaurant) error {
	return uc.repo.CreateRestaurant(ctx, entity.Restaurant{
		Id:       uuid.New(),
		Name:     req.Name,
		Location: req.Location,
	})
}
