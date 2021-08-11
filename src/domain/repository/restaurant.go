package repository

import "restaurant-record/domain/model"

type RestaurantRepository interface {
	FindNear(location string) ([]*model.Restaurant, error)
}
