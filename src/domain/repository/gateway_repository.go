package repository

import "gorm.io/gorm"

type GatewayRepository struct {
	db *gorm.DB
}

func NewGatewayRepository(db *gorm.DB) *GatewayRepository {
	return &GatewayRepository{
		db: db,
	}
}

type IGatewayRepository interface {
	GetAll()
	Get()
	Save()
	Update()
}

func (r *GatewayRepository) GetAll() {

}

func (r *GatewayRepository) Get() {

}

func (r *GatewayRepository) Save() {

}

func (r *GatewayRepository) Update() {

}
