package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	key := ""
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s&key=%s&radius=200&language=ja&type=restaurant", location, key)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var data Response

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
