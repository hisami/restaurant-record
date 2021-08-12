package handler

import (
	"net/http"
	"restaurant-record/usecase"

	"github.com/gin-gonic/gin"
)

type RestaurantPostRequest struct {
	Name string `json:"name" binding:"required"`
}

func RestaurantCreate(ru usecase.RestaurantUsecase) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var req RestaurantPostRequest
		// バリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		restaurant, _ := ru.Create(req.Name)
		c.JSON(http.StatusOK, restaurant)
	}
	return gin.HandlerFunc(fn)
}
