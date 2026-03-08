package rest

import (
	"bcc-workshop-2/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, uc *usecase.Usecase) {
	restaurantCtrl := NewRestaurantController(uc.RestaurantUsecase)

	api := router.Group("/api/v1")
	{
		restaurant := api.Group("/restaurants")
		{
			restaurant.POST("", restaurantCtrl.CreateRestaurant)
			restaurant.GET("", restaurantCtrl.GetAllRestaurants)
		}
	}
}
