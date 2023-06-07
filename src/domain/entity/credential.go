package entity

type Credential struct {
	ID          int64    `gorm:"primaryKey" json:"id"`
	GatewayID   int64    `json:"-"`
	Gateway     *Gateway `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gateway,omitempty"`
	Url         string   `gorm:"size:200" json:"url"`
	MerchantId  string   `gorm:"size:65" json:"merchant_id"`
	Password    string   `gorm:"size:65" json:"password"`
	MerchantKey string   `gorm:"size:65" json:"merchant_key"`
	SecretKey   string   `gorm:"size:65" json:"secret_key"`
}

func (e *Credential) GetUrl() string {
	return e.Url
}

func (e *Credential) GetMerchantId() string {
	return e.MerchantId
}

func (e *Credential) GetPassword() string {
	return e.Password
}

func (e *Credential) GetMerchantKey() string {
	return e.MerchantKey
}

func (e *Credential) GetSecretKey() string {
	return e.SecretKey
}
