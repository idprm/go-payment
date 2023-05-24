package entity

type Application struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	Code        string `gorm:"size:20" json:"code"`
	Name        string `gorm:"size:60" json:"name"`
	Url         string `gorm:"size:50" json:"url"`
	UrlCallback string `gorm:"size:150" json:"url_callback"`
}

func (e *Application) GetId() int64 {
	return e.ID
}

func (e *Application) GetCode() string {
	return e.Code
}

func (e *Application) GetName() string {
	return e.Name
}

func (e *Application) GetUrl() string {
	return e.Url
}

func (e *Application) GetUrlCallback() string {
	return e.UrlCallback
}
