package entity

import "gorm.io/gorm"

type Return struct {
	ID         int64    `gorm:"primaryKey" json:"id"`
	PaymentID  int64    `json:"-"`
	Payment    *Payment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payment,omitempty"`
	Payload    string   `gorm:"type:text" json:"payload"`
	gorm.Model `json:"model"`
}

func (e *Return) GetId() int64 {
	return e.ID
}

func (e *Return) GetPayload() string {
	return e.Payload
}
