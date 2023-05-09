package jazzcash

import "github.com/idprm/go-payment/src/config"

type JazzCash struct {
	conf *config.Secret
}

func NewJazzCash(conf *config.Secret) *JazzCash {
	return &JazzCash{
		conf: conf,
	}
}

func (p *JazzCash) Payment() {

}
