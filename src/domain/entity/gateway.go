package entity

type Gateway struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Code string `gorm:"size:20" json:"code"`
	Name string `gorm:"size:60" json:"name"`
}
