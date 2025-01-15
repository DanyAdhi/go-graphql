package orders

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll(requestFields []string) (*[]Orders, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r repository) GetAll(requestFields []string) (*[]Orders, error) {
	var orders []Orders
	err := r.db.Select(requestFields).Find(&orders).Error
	if err != nil {
		log.Printf("Error get pets. %v", err)
		return nil, err
	}

	return &orders, nil
}
