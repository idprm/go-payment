package callback

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
)

type Callback struct {
	conf   *config.Secret
	logger *logger.Logger
	app    *entity.Application
	order  *entity.Order
}

func NewCallback(
	conf *config.Secret,
	logger *logger.Logger,
	app *entity.Application,
	order *entity.Order,
) *Callback {
	return &Callback{
		conf:   conf,
		logger: logger,
		app:    app,
		order:  order,
	}
}

func (p *Callback) Hit() ([]byte, error) {
	request := &entity.CallbackRequestBody{
		Number: p.order.GetNumber(),
		IsPaid: true,
		Time:   time.Now(),
	}
	payload, _ := json.Marshal(&request)
	req, err := http.NewRequest("POST", p.app.GetUrlCallback(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		Proxy:              http.ProxyFromEnvironment,
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	p.logger.Writer(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	p.logger.Writer(body)
	return body, nil
}
