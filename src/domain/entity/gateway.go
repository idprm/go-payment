package entity

type Gateway struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	CountryID int64      `json:"-"`
	Country   *Country   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	Code      string     `gorm:"size:20" json:"code"`
	Name      string     `gorm:"size:60" json:"name"`
	Channel   *[]Channel `gorm:"foreignKey:gateway_id" json:"channels,omitempty"`
}
