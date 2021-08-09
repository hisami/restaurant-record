package main

import (
	"net/http"
	"restaurant-record/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	restaurantUsecase := usecase.NewRestaurantUsecase()

	// ルーティング
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		location := c.Query("location")
		restaurantUsecase.FindNear(location)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	router.Run()
}
