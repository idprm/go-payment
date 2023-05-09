package repository

import "gorm.io/gorm"

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

type IOrderRepository interface {
	GetAll()
	Get()
	Save()
	Update()
}

func (r *OrderRepository) GetAll() {

}

func (r *OrderRepository) Get() {

}

func (r *OrderRepository) Save() {

}

func (r *OrderRepository) Update() {

}
