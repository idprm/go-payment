package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type CallbackRepository struct {
	db *gorm.DB
}

func NewCallbackRepository(db *gorm.DB) *CallbackRepository {
	return &CallbackRepository{
		db: db,
	}
}

type ICallbackRepository interface {
	GetAll() (*[]entity.Callback, error)
	GetById(int) (*entity.Callback, error)
	Save(*entity.Callback) (*entity.Callback, error)
	Update(*entity.Callback) (*entity.Callback, error)
	Delete(int) error
}

func (r *CallbackRepository) GetAll() (*[]entity.Callback, error) {
	var callbacks []entity.Callback
	err := r.db.Order("id desc").Preload("Order").Find(&callbacks).Error
	if err != nil {
		return nil, err
	}
	return &callbacks, err
}

func (r *CallbackRepository) GetById(id int) (*entity.Callback, error) {
	var callback entity.Callback
	err := r.db.Where("id = ?", id).Preload("Order").Take(&callback).Error
	if err != nil {
		return nil, err
	}
	return &callback, err
}

func (r *CallbackRepository) Save(callback *entity.Callback) (*entity.Callback, error) {
	err := r.db.Create(&callback).Error
	if err != nil {
		return nil, err
	}
	return callback, nil
}

func (r *CallbackRepository) Update(callback *entity.Callback) (*entity.Callback, error) {
	err := r.db.Save(&callback).Error
	if err != nil {
		return nil, err
	}
	return callback, nil
}

func (r *CallbackRepository) Delete(id int) error {
	var callback entity.Callback
	err := r.db.Where("id = ?", id).Delete(&callback).Error
	if err != nil {
		return err
	}
	return nil
}
