package entity

type Country struct {
	ID      int        `gorm:"primaryKey" json:"id"`
	Name    string     `gorm:"size:50;not null" json:"name"`
	Locale  string     `gorm:"size:15" json:"locale"`
	Prefix  string     `gorm:"size:3" json:"prefix"`
	Flag    string     `gorm:"size:50" json:"flag"`
	Gateway *[]Gateway `gorm:"foreignKey:country_id" json:"gateways,omitempty"`
}

func (e *Country) GetId() int {
	return e.ID
}

func (e *Country) GetName() string {
	return e.Name
}

func (e *Country) GetLocale() string {
	return e.Locale
}

func (e *Country) GetPrefix() string {
	return e.Prefix
}

func (e *Country) GetFlag() string {
	return e.Flag
}
