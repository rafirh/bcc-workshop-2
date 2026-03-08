package rest

import (
	"net/http"
	"strconv"

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

func (rc *RestaurantController) GetAllRestaurants(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	pagination := model.Pagination{Page: page, Limit: limit}

	restaurants, err := rc.usecase.GetAllRestaurants(c.Request.Context(), pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": restaurants,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
		},
	})
}
