package main

import (
	"fmt"
	"log"
	"os"
	"restaurant-record/infra"
	"restaurant-record/presentation/handler"
	"restaurant-record/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

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
	api := router.Group("/api")
	api.GET("/google-restaurants", handler.GooogleRestaurantCreate(googleRestaurantUsecase))
	api.POST("/restaurants", handler.RestaurantCreate(restaurantUsecase))

	router.Run()
}
