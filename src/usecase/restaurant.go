package usecase

import (
	"context"
	"log"
	"os"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

type RestaurantUsecase interface {
	FindNear(location string)
}

type restaurantUsecase struct{}

// コンストラクタ
func NewRestaurantUsecase() RestaurantUsecase {
	return &restaurantUsecase{}
}

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
}

// 近隣の飲食店を返却
func (ru *restaurantUsecase) FindNear(location string) {
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
	pretty.Println(res.Results)

	// 最終的にはドメインオブジェクトの配列にしたい

	// key := os.Getenv("API_KEY")
	// url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s&key=%s&radius=200&language=ja&type=restaurant", location, key)
	// resp, err := http.Get(url)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// body, err := ioutil.ReadAll(resp.Body)

	// var data Response

	// if err := json.Unmarshal(body, &data); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(data)
}
