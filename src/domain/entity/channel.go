package entity

type Channel struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	GatewayID int64 `json:"gateway_id"`
	Gateway   *Gateway
	Name      string `gorm:"size:80" json:"name"`
	Slug      string `gorm:"size:60" json:"slug"`
	Logo      string `gorm:"size:60" json:"logo"`
	Type      string `gorm:"size:60" json:"type"`
	Param     string `gorm:"size:60" json:"param"`
	IsActive  bool   `gorm:"type:bool" json:"is_active"`
}

func (e *Channel) GetId() int64 {
	return e.ID
}

func (e *Channel) GetName() string {
	return e.Name
}
