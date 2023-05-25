package entity

import "gorm.io/gorm"

type Callback struct {
	ID         int64    `gorm:"primaryKey" json:"id"`
	PaymentID  int64    `json:"payment_id"`
	Payment    *Payment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payment,omitempty"`
	Payload    string   `gorm:"type:text" json:"payload"`
	gorm.Model `json:"-"`
}
