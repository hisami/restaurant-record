package main

import (
	"log"
	"net/http"
	"restaurant-record/infra"
	"restaurant-record/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DI
	googleRestaurantRepository := infra.NewGoogleRestaurantRepository()
	googleRestaurantUsecase := usecase.NewGoogleRestaurantUsecase(googleRestaurantRepository)

	// ルーティング
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		location := c.Query("location")
		restaurants, _ := googleRestaurantUsecase.FindNear(location)
		c.JSON(http.StatusOK, restaurants)
	})

	router.Run()
}
