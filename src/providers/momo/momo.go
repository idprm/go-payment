package momo

import "github.com/idprm/go-payment/src/config"

type Momo struct {
	conf *config.Secret
}

func NewMomo(conf *config.Secret) *Momo {
	return &Momo{
		conf: conf,
	}
}

func (p *Momo) Payment() {

}
