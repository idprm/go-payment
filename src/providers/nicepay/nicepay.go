package nicepay

import "github.com/idprm/go-payment/src/config"

type Nicepay struct {
	conf *config.Secret
}

func NewNicepay(conf *config.Secret) *Nicepay {
	return &Nicepay{
		conf: conf,
	}
}

func (p *Nicepay) Payment() {

}
