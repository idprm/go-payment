package entity

type Country struct {
	ID      int        `gorm:"primaryKey" json:"id"`
	Name    string     `gorm:"size:50;not null" json:"name"`
	Locale  string     `gorm:"size:15" json:"locale"`
	Prefix  string     `gorm:"size:3" json:"prefix"`
	Flag    string     `gorm:"size:50" json:"flag"`
	Gateway *[]Gateway `gorm:"foreignKey:country_id" json:"gateways,omitempty"`
}
