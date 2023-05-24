package entity

import "gorm.io/gorm"

type Payment struct {
	ID         int64 `gorm:"primaryKey" json:"id"`
	OrderID    int64 `json:"order_id"`
	Order      *Order
	IsPaid     bool `gorm:"type:bool" json:"is_paid"`
	gorm.Model `json:"model"`
}
