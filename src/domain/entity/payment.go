package entity

import "gorm.io/gorm"

type Payment struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	OrderID    int64  `json:"order_id"`
	Order      *Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order,omitempty"`
	IsPaid     bool   `gorm:"type:bool" json:"is_paid"`
	gorm.Model `json:"model"`
}
