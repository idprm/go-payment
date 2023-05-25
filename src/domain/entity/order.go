package entity

import "gorm.io/gorm"

type Order struct {
	ID            int64        `gorm:"primaryKey" json:"id"`
	ApplicationID int64        `json:"application_id"`
	Application   *Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"application,omitempty"`
	ChannelID     int64        `json:"channel_id"`
	Channel       *Channel     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"channel,omitempty"`
	Number        string       `gorm:"size:56" json:"number"`
	Msisdn        string       `gorm:"size:25" json:"msisdn"`
	Email         string       `gorm:"size:100" json:"email"`
	Name          string       `gorm:"size:150" json:"name"`
	Amount        float64      `gorm:"size:15" json:"amount"`
	Description   string       `gorm:"size:100" json:"description"`
	IpAddress     string       `gorm:"size:25" json:"ip_address"`
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

func (e *Order) GetName() string {
	return e.Name
}

func (e *Order) GetDescription() string {
	return e.Description
}
