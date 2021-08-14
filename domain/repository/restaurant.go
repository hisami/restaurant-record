package repository

import "restaurant-record/domain/model"

type RestaurantRepository interface {
	Create(restaurant *model.Restaurant) (*model.Restaurant, error)
	FindAll() ([]*model.Restaurant, error)
}
