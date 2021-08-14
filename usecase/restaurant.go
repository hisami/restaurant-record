package usecase

import (
	"restaurant-record/domain/model"
	"restaurant-record/domain/repository"
)

type RestaurantUsecase interface {
	Create(name string) (*model.Restaurant, error)
	FindAll() ([]*model.Restaurant, error)
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

// Create
func (ru *restaurantUsecase) Create(name string) (*model.Restaurant, error) {
	restaurant, err := model.NewRestaurant(name)
	if err != nil {
		return nil, err
	}

	createdRestaurant, err := ru.restaurantRepo.Create(restaurant)
	if err != nil {
		return nil, err
	}

	return createdRestaurant, nil
}

// FindAll
func (ru *restaurantUsecase) FindAll() ([]*model.Restaurant, error) {
	restaurants, err := ru.restaurantRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return restaurants, nil
}
