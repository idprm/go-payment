package repository

import "gorm.io/gorm"

type RefundRepository struct {
	db *gorm.DB
}

func NewRefundRepository(db *gorm.DB) *RefundRepository {
	return &RefundRepository{
		db: db,
	}
}

type IRefundRepository interface {
	GetAll()
	Get()
	Save()
	Update()
}

func (r *RefundRepository) GetAll() {

}

func (r *RefundRepository) Get() {

}

func (r *RefundRepository) Save() {

}

func (r *RefundRepository) Update() {

}
