package infra

import (
	"restaurant-record/domain/model"
	"restaurant-record/domain/repository"

	"github.com/jinzhu/gorm"
)

type RestaurantRepository struct {
	Conn *gorm.DB
}

// コンストラクタ
func NewRestaurantRepository(conn *gorm.DB) repository.RestaurantRepository {
	return &RestaurantRepository{Conn: conn}
}

// Create
func (rr *RestaurantRepository) Create(restaurant *model.Restaurant) (*model.Restaurant, error) {
	if err := rr.Conn.Create(&restaurant).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}

// FindAll
func (rr *RestaurantRepository) FindAll() ([]*model.Restaurant, error) {
	var restaurants []*model.Restaurant

	if err := rr.Conn.Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}
