package entity

import "gorm.io/gorm"

type Order struct {
	ID        int64   `gorm:"primaryKey" json:"id"`
	Number    string  `gorm:"size:56" json:"number"`
	Msisdn    string  `gorm:"size:25" json:"msisdn"`
	Email     string  `gorm:"size:56" json:"email"`
	Amount    float64 `gorm:"size:15" json:"amount"`
	IpAddress string  `gorm:"size:25" json:"ip_address"`
	gorm.Model
}

func (e *Order) GetNumber() string {
	return e.Number
}

func (e *Order) GetMsisdn() string {
	return e.Msisdn
}

func (e *Order) GetEmail() string {
	return e.Email
}
