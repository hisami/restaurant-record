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

type RestaurantRequest struct {
	Name string `json:"name"`
}

func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB接続
	connect := fmt.Sprintf("%s:%s@%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	fmt.Println(connect)
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic("Cannot open database")
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
		location := c.Query("location")
		restaurants, _ := googleRestaurantUsecase.FindNear(location)
		c.JSON(http.StatusOK, restaurants)
	})

	router.POST("/", func(c *gin.Context) {
		// リクエストボディの受け取り
		var restaurantRequest RestaurantRequest
		c.BindJSON(&restaurantRequest)
		restaurant, _ := restaurantUsecase.Create(restaurantRequest.Name)
		c.JSON(http.StatusOK, restaurant)
	})

	router.Run()
}
