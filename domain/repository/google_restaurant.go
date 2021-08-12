package repository

import "restaurant-record/domain/model"

type GoogleRestaurantRepository interface {
	FindNear(location string) ([]*model.GoogleRestaurant, error)
}
