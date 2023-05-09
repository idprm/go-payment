package entity

import "gorm.io/gorm"

type Callback struct {
	ID         int64 `gorm:"primaryKey" json:"id"`
	gorm.Model `json:"-"`
}
