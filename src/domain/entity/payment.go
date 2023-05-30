package entity

import "gorm.io/gorm"

type Payment struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	OrderID    int64     `json:"-"`
	Order      *Order    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order,omitempty"`
	IsPaid     bool      `gorm:"type:bool" json:"is_paid"`
	Callback   *Callback `gorm:"foreignKey:payment_id" json:"callback,omitempty"`
	gorm.Model `json:"model"`
}

func (e *Payment) GetId() int64 {
	return e.ID
}

func (e *Payment) GetIsPaid() bool {
	return e.IsPaid
}
