package entity

import "gorm.io/gorm"

type Transaction struct {
	ID            int64 `gorm:"primaryKey" json:"id"`
	ApplicationID int64 `json:"application_id"`
	Application   *Application
	Action        string `gorm:"size:150" json:"action"`
	Payload       string `gorm:"type:text" json:"payload"`
	gorm.Model
}
