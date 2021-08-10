package usecase

import (
	"context"
	"log"
	"os"
	"restaurant-record/domain/model"
	"strconv"
	"strings"

	"googlemaps.github.io/maps"
)

type RestaurantUsecase interface {
	FindNear(location string) ([]*model.Restaurant, error)
}

type restaurantUsecase struct{}

// コンストラクタ
func NewRestaurantUsecase() RestaurantUsecase {
	return &restaurantUsecase{}
}

// 近隣の飲食店を返却
func (ru *restaurantUsecase) FindNear(location string) ([]*model.Restaurant, error) {
	// クライアントの生成
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// クエリの生成
	splittedLocation := strings.Split(location, ",")
	lat, _ := strconv.ParseFloat(splittedLocation[0], 64)
	lng, _ := strconv.ParseFloat(splittedLocation[1], 64)
	mapLocation := maps.LatLng{Lat: lat, Lng: lng}
	r := &maps.NearbySearchRequest{
		Location: &mapLocation,
		Radius:   200,
		Type:     "restaurant",
		Language: "ja",
	}

	// リクエスト送信
	res, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// 結果をドメインオブジェクトの配列に詰め替え
	var restaurants []*model.Restaurant
	for _, v := range res.Results {
		restaurants = append(restaurants, &model.Restaurant{
			Name: v.Name,
		})
	}

	return restaurants, nil
}
