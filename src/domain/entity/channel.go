package entity

type Channel struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	GatewayID int64    `json:"-"`
	Gateway   *Gateway `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gateway,omitempty"`
	Name      string   `gorm:"size:80" json:"name"`
	Slug      string   `gorm:"size:60" json:"slug"`
	Logo      string   `gorm:"size:60" json:"logo"`
	Type      string   `gorm:"size:60" json:"type"`
	Param     string   `gorm:"size:60" json:"param"`
	IsActive  bool     `gorm:"type:bool" json:"is_active"`
}

func (e *Channel) GetId() int64 {
	return e.ID
}

func (e *Channel) GetName() string {
	return e.Name
}

func (e *Channel) GetSlug() string {
	return e.Slug
}

func (e *Channel) GetLogo() string {
	return e.Logo
}

func (e *Channel) SetLogo(url, logo string) {
	e.Logo = url + "/static/images/payment/" + logo
}

func (e *Channel) GetType() string {
	return e.Type
}

func (e *Channel) GetParam() string {
	return e.Param
}

func (e *Channel) GetIsActive() bool {
	return e.IsActive
}
