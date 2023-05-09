package repository

import "gorm.io/gorm"

type CallbackRepository struct {
	db *gorm.DB
}

func NewCallbackRepository(db *gorm.DB) *CallbackRepository {
	return &CallbackRepository{
		db: db,
	}
}

type ICallbackRepository interface {
	GetAll()
	Get()
	Save()
	Update()
}

func (r *CallbackRepository) GetAll() {

}

func (r *CallbackRepository) Get() {

}

func (r *CallbackRepository) Save() {

}

func (r *CallbackRepository) Update() {

}
