package usecase

import (
	"restaurant-record/domain/model"
	"restaurant-record/domain/repository"
)

type GoogleRestaurantUsecase interface {
	FindNear(location string) ([]*model.GoogleRestaurant, error)
}

type googleRestaurantUsecase struct {
	googleRestaurantRepo repository.GoogleRestaurantRepository
}

// コンストラクタ
func NewGoogleRestaurantUsecase(googleRestaurantRepo repository.GoogleRestaurantRepository) GoogleRestaurantUsecase {
	return &googleRestaurantUsecase{
		googleRestaurantRepo: googleRestaurantRepo,
	}
}

// 近隣の飲食店を返却
func (ru *googleRestaurantUsecase) FindNear(location string) ([]*model.GoogleRestaurant, error) {
	googleRestaurants, err := ru.googleRestaurantRepo.FindNear(location)
	if err != nil {
		return nil, err
	}

	return googleRestaurants, nil
}
