package midtrans

import "github.com/idprm/go-payment/src/config"

type Midtrans struct {
	conf *config.Secret
}

func NewMidtrans(conf *config.Secret) *Midtrans {
	return &Midtrans{
		conf: conf,
	}
}

func (p *Midtrans) Payment() {

}
