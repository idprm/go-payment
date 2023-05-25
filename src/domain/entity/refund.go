package entity

import "gorm.io/gorm"

type Refund struct {
	ID            int64        `gorm:"primaryKey" json:"id"`
	ApplicationID int64        `json:"application_id"`
	Application   *Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"application,omitempty"`
	ChannelID     int64        `json:"channel_id"`
	Channel       *Channel     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"channel,omitempty"`
	Number        string       `gorm:"size:56" json:"number"`
	Amount        float64      `gorm:"size:15" json:"amount"`
	IpAddress     string       `gorm:"size:25" json:"ip_address"`
	gorm.Model
}

func (e *Refund) GetNumber() string {
	return e.Number
}

func (e *Refund) GetAmount() float64 {
	return e.Amount
}
