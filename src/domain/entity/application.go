package entity

type Application struct {
	ID          int64          `gorm:"primaryKey" json:"id"`
	UrlCallback string         `gorm:"size:200" json:"url_callback"`
	Code        string         `gorm:"size:20" json:"code"`
	Name        string         `gorm:"size:60" json:"name"`
	Order       *[]Order       `gorm:"foreignKey:application_id" json:"transaction,omitempty"`
	Transaction *[]Transaction `gorm:"foreignKey:application_id" json:"callback,omitempty"`
}

func (e *Application) GetId() int64 {
	return e.ID
}

func (e *Application) GetUrlCallback() string {
	return e.UrlCallback
}

func (e *Application) GetCode() string {
	return e.Code
}

func (e *Application) GetName() string {
	return e.Name
}

func (e *Application) IsSehatCepat() bool {
	return e.Code == "sehatcepat"
}

func (e *Application) IsSuratSakit() bool {
	return e.Code == "suratsakit"
}

func (e *Application) IsSurkit() bool {
	return e.Code == "suratsakit_new"
}
