package handler

import (
	"net/http"
	"restaurant-record/usecase"

	"github.com/gin-gonic/gin"
)

type GoogleRestaurantGetRequest struct {
	Location string `form:"location" binding:"required"`
}

func GooogleRestaurantCreate(ru usecase.GoogleRestaurantUsecase) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var req GoogleRestaurantGetRequest
		// バリデーション
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		restaurants, _ := ru.FindNear(req.Location)
		c.JSON(http.StatusOK, restaurants)
	}
	return gin.HandlerFunc(fn)
}
