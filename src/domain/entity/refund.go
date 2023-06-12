package entity

import "gorm.io/gorm"

type Refund struct {
	ID         int64    `gorm:"primaryKey" json:"id"`
	PaymentID  int64    `json:"-"`
	Payment    *Payment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payment,omitempty"`
	Number     string   `gorm:"size:56" json:"number"`
	Amount     float64  `gorm:"size:15" json:"amount"`
	Status     string   `gorm:"size:55" json:"status"`
	IpAddress  string   `gorm:"size:25" json:"ip_address"`
	gorm.Model `json:"model"`
}

func (e *Refund) GetNumber() string {
	return e.Number
}

func (e *Refund) GetAmount() float64 {
	return e.Amount
}

func (e *Refund) GetStatus() string {
	return e.Status
}
