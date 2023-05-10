package entity

import "gorm.io/gorm"

type Transaction struct {
	ID      int64 `gorm:"primaryKey" json:"id"`
	OrderID int64 `json:"order_id"`
	Order   *Order
	Payload string `gorm:"type:text" json:"payload"`
	gorm.Model
}
