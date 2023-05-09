package razer

import "github.com/idprm/go-payment/src/config"

type Razer struct {
	conf *config.Secret
}

func NewRazer(conf *config.Secret) *Razer {
	return &Razer{
		conf: conf,
	}
}

func (p *Razer) Payment() {

}
