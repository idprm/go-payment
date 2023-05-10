package momo

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
)

type Momo struct {
	conf  *config.Secret
	order *entity.Order
}

func NewMomo(
	conf *config.Secret,
	order *entity.Order,
) *Momo {
	return &Momo{
		conf:  conf,
		order: order,
	}
}

func (p *Momo) Payment() {

}
