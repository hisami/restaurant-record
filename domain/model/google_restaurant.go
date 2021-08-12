package model

type GoogleRestaurant struct {
	Name string `json:"name"`
}

// コンストラクタ
func NewGoogleRestaurant(name string) (*GoogleRestaurant, error) {
	googleRestaurant := &GoogleRestaurant{
		Name: name,
	}

	return googleRestaurant, nil
}
