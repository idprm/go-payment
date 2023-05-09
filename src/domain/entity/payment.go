package entity

import "gorm.io/gorm"

type Payment struct {
	ID      int64 `gorm:"primaryKey" json:"id"`
	OrderID int64 `json:"order_id"`
	Order   *Order
	gorm.Model
}
