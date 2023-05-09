package entity

type Application struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	Code        string `gorm:"size:20" json:"code"`
	Name        string `gorm:"size:60" json:"name"`
	UrlCallback string `gorm:"size:150" json:"url_callback"`
}
