package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

type IOrderRepository interface {
	GetAll() (*[]entity.Order, error)
	GetById(int) (*entity.Order, error)
	GetByNumber(int, string) (*entity.Order, error)
	CountByNumber(int, string) (int64, error)
	Save(*entity.Order) (*entity.Order, error)
	Update(*entity.Order) (*entity.Order, error)
	Delete(int) error
}

func (r *OrderRepository) GetAll() (*[]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Order("id desc").Preload("Application").Preload("Channel").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, err
}

func (r *OrderRepository) GetById(id int) (*entity.Order, error) {
	var order entity.Order
	err := r.db.Where("id = ?", id).Preload("Application").Preload("Channel").Take(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, err
}

func (r *OrderRepository) GetByNumber(appId int, number string) (*entity.Order, error) {
	var order entity.Order
	err := r.db.Where("application_id = ?", appId).Where("number = ?", number).Preload("Application").Preload("Channel").Take(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, err
}

func (r *OrderRepository) CountByNumber(appId int, number string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Order{}).Where("application_id = ?", appId).Where("number = ?", number).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *OrderRepository) Save(order *entity.Order) (*entity.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) Update(order *entity.Order) (*entity.Order, error) {
	err := r.db.Save(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) Delete(id int) error {
	var order entity.Order
	err := r.db.Where("id = ?", id).Delete(&order).Error
	if err != nil {
		return err
	}
	return nil
}
