package usecase

import (
	"context"
	"log"
	"os"
	"restaurant-record/domain/model"

	"googlemaps.github.io/maps"
)

type RestaurantUsecase interface {
	FindNear(location string) (*[]model.Restaurant, error)
}

type restaurantUsecase struct{}

// コンストラクタ
func NewRestaurantUsecase() RestaurantUsecase {
	return &restaurantUsecase{}
}

// 近隣の飲食店を返却
func (ru *restaurantUsecase) FindNear(location string) (*[]model.Restaurant, error) {
	mapLocation := maps.LatLng{Lat: 34.39339788681548, Lng: 132.42305207662446}

	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.NearbySearchRequest{
		Location: &mapLocation,
		Radius:   200,
		Type:     "restaurant",
		Language: "ja",
	}

	res, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// ドメインオブジェクトの配列に詰め替え
	var restaurants []model.Restaurant

	for _, v := range res.Results {
		restaurants = append(restaurants, model.Restaurant{
			Name: v.Name,
		})
	}

	return &restaurants, nil
}
