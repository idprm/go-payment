package nicepay

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
)

type Nicepay struct {
	conf  *config.Secret
	order *entity.Order
}

func NewNicepay(
	conf *config.Secret,
	order *entity.Order,
) *Nicepay {
	return &Nicepay{
		conf:  conf,
		order: order,
	}
}

func (p *Nicepay) Payment() {

}
