package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type GatewayRepository struct {
	db *gorm.DB
}

func NewGatewayRepository(db *gorm.DB) *GatewayRepository {
	return &GatewayRepository{
		db: db,
	}
}

type IGatewayRepository interface {
	GetAll() (*[]entity.Gateway, error)
	GetByCode(string) (*entity.Gateway, error)
	Save(*entity.Gateway) (*entity.Gateway, error)
	Update(*entity.Gateway) (*entity.Gateway, error)
	Delete(int) error
}

func (r *GatewayRepository) GetAll() (*[]entity.Gateway, error) {
	var gateways []entity.Gateway
	err := r.db.Order("id asc").Preload("Channel").Preload("Country").Find(&gateways).Error
	if err != nil {
		return nil, err
	}
	return &gateways, err
}

func (r *GatewayRepository) GetByCode(code string) (*entity.Gateway, error) {
	var gateway entity.Gateway
	err := r.db.Where("code = ?", code).Preload("Channel").Preload("Country").Take(&gateway).Error
	if err != nil {
		return nil, err
	}
	return &gateway, err
}

func (r *GatewayRepository) Save(gateway *entity.Gateway) (*entity.Gateway, error) {
	err := r.db.Create(&gateway).Error
	if err != nil {
		return nil, err
	}
	return gateway, nil
}

func (r *GatewayRepository) Update(gateway *entity.Gateway) (*entity.Gateway, error) {
	err := r.db.Save(&gateway).Error
	if err != nil {
		return nil, err
	}
	return gateway, nil
}

func (r *GatewayRepository) Delete(id int) error {
	var gateway entity.Gateway
	err := r.db.Where("id = ?", id).Delete(&gateway).Error
	if err != nil {
		return err
	}
	return nil
}
