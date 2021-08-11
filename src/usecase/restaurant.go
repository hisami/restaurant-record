package usecase

import (
	"restaurant-record/domain/model"
	"restaurant-record/domain/repository"
)

type RestaurantUsecase interface {
	FindNear(location string) ([]*model.Restaurant, error)
}

type restaurantUsecase struct {
	restaurantRepo repository.RestaurantRepository
}

// コンストラクタ
func NewRestaurantUsecase(restaurantRepo repository.RestaurantRepository) RestaurantUsecase {
	return &restaurantUsecase{
		restaurantRepo: restaurantRepo,
	}
}

// 近隣の飲食店を返却
func (ru *restaurantUsecase) FindNear(location string) ([]*model.Restaurant, error) {
	restaurants, err := ru.restaurantRepo.FindNear(location)
	if err != nil {
		return nil, err
	}

	return restaurants, nil
}
