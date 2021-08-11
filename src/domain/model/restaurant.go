package model

type Restaurant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// コンストラクタ
func NewRestaurant(name string) (*Restaurant, error) {
	restaurant := &Restaurant{
		Name: name,
	}

	return restaurant, nil
}
