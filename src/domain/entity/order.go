package entity

import (
	"strings"

	"gorm.io/gorm"
)

type Order struct {
	ID            int64        `gorm:"primaryKey" json:"id"`
	ApplicationID int64        `json:"-"`
	Application   *Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"application,omitempty"`
	ChannelID     int64        `json:"-"`
	Channel       *Channel     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"channel,omitempty"`
	UrlReturn     string       `gorm:"size:200" json:"url_return"`
	Number        string       `gorm:"size:60;uniqueIndex:uidx_number" json:"number"`
	Msisdn        string       `gorm:"size:25" json:"msisdn"`
	Email         string       `gorm:"size:100" json:"email"`
	Name          string       `gorm:"size:150" json:"name"`
	Amount        float64      `gorm:"size:15" json:"amount"`
	Description   string       `gorm:"size:100" json:"description"`
	IpAddress     string       `gorm:"size:25" json:"ip_address"`
	gorm.Model
}

func (e *Order) GetId() int64 {
	return e.ID
}

func (e *Order) GetUrlReturn() string {
	return e.UrlReturn
}

func (e *Order) GetNumber() string {
	return e.Number
}

func (e *Order) GetMsisdn() string {
	return e.Msisdn
}

func (e *Order) SetMsisdn() {
	// remove (+) character
	r := strings.NewReplacer("+", "")
	e.Msisdn = r.Replace(e.Msisdn)
}

func (e *Order) GetEmail() string {
	return e.Email
}

func (e *Order) GetName() string {
	return e.Name
}

func (e *Order) GetAmount() float64 {
	return e.Amount
}

func (e *Order) GetDescription() string {
	return e.Description
}

func (e *Order) GetIpAddress() string {
	return e.IpAddress
}
