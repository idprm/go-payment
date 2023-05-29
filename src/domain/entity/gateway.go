package entity

type Gateway struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	CountryID int64      `json:"-"`
	Country   *Country   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	Code      string     `gorm:"size:20" json:"code"`
	Name      string     `gorm:"size:60" json:"name"`
	Currency  string     `gorm:"size:15" json:"currency"`
	Channel   *[]Channel `gorm:"foreignKey:gateway_id" json:"channels,omitempty"`
}

func (e *Gateway) GetId() int64 {
	return e.ID
}

func (e *Gateway) GetCode() string {
	return e.Code
}

func (e *Gateway) GetName() string {
	return e.Name
}

func (e *Gateway) GetCurrency() string {
	return e.Currency
}
