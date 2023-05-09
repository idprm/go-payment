package postback

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
)

type Callback struct {
	conf *config.Secret
	app  *entity.Application
}

func NewCallback(
	conf *config.Secret,
	app *entity.Application,
) *Callback {
	return &Callback{
		conf: conf,
		app:  app,
	}
}

func (p *Callback) Hit() {

}
