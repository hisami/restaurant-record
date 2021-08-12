package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restaurant-record/infra"
	"restaurant-record/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type GoogleRestaurantGetRequest struct {
	Location string `form:"location" binding:"required"`
}

type RestaurantPostRequest struct {
	Name string `json:"name" binding:"required"`
}

func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot not find .env file")
	}

	// DB接続
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		log.Println("Cannot open database")
	}
	defer db.Close()

	// DI
	googleRestaurantRepository := infra.NewGoogleRestaurantRepository()
	googleRestaurantUsecase := usecase.NewGoogleRestaurantUsecase(googleRestaurantRepository)
	restaurantRepository := infra.NewRestaurantRepository(db)
	restaurantUsecase := usecase.NewRestaurantUsecase(restaurantRepository)

	// ルーティング
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		var req GoogleRestaurantGetRequest
		// バリデーション
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		restaurants, _ := googleRestaurantUsecase.FindNear(req.Location)
		c.JSON(http.StatusOK, restaurants)
	})

	router.POST("/", func(c *gin.Context) {
		var req RestaurantPostRequest
		// バリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		restaurant, _ := restaurantUsecase.Create(req.Name)
		c.JSON(http.StatusOK, restaurant)
	})

	router.Run()
}
