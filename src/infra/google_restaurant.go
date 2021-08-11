package infra

import (
	"context"
	"log"
	"os"
	"restaurant-record/domain/model"
	"restaurant-record/domain/repository"
	"strconv"
	"strings"

	"googlemaps.github.io/maps"
)

type RestaurantRepository struct{}

// コンストラクタ
func NewGoogleRestaurantRepository() repository.GoogleRestaurantRepository {
	return &RestaurantRepository{}
}

// Google APIを使って近隣の飲食店を返却
func (rr *RestaurantRepository) FindNear(location string) ([]*model.GoogleRestaurant, error) {
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
	var googleRestaurants = make([]*model.GoogleRestaurant, 0)
	for _, v := range res.Results {
		googleRestaurants = append(googleRestaurants, &model.GoogleRestaurant{
			Name: v.Name,
		})
	}

	return googleRestaurants, nil
}
