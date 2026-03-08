package usecase

import "bcc-workshop-2/internal/repository"

type Usecase struct {
	RestaurantUsecase IRestaurantUsecase
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		RestaurantUsecase: NewRestaurantUseCase(repo.RestaurantRepository),
	}
}
