package razer

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
)

type Razer struct {
	conf  *config.Secret
	order *entity.Order
}

func NewRazer(
	conf *config.Secret,
	order *entity.Order,
) *Razer {
	return &Razer{
		conf:  conf,
		order: order,
	}
}

func (p *Razer) Payment() {

}
