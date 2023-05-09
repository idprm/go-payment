package entity

type Channel struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:60" json:"name"`
	Param    string `gorm:"size:60" json:"param"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
}
