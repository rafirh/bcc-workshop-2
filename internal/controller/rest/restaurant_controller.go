package rest

import (
	"net/http"

	"bcc-workshop-2/internal/model"
	"bcc-workshop-2/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RestaurantController struct {
	usecase usecase.IRestaurantUsecase
}

func NewRestaurantController(uc usecase.IRestaurantUsecase) *RestaurantController {
	return &RestaurantController{usecase: uc}
}

func (rc *RestaurantController) CreateRestaurant(c *gin.Context) {
	var req model.CreateRestaurant
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := rc.usecase.CreateRestaurant(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "restaurant created successfully",
	})
}
