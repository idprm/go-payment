package dragonpay

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

type DragonPay struct {
	conf   *config.Secret
	logger *logger.Logger
	order  *entity.Order
}

func NewDragonPay(
	conf *config.Secret,
	logger *logger.Logger,
	order *entity.Order,
) *DragonPay {
	return &DragonPay{
		conf:   conf,
		logger: logger,
		order:  order,
	}
}

func (p *DragonPay) Payment() ([]byte, error) {
	url := p.conf.DragonPay.Url + p.order.GetNumber() + "/post"
	var request entity.DragonPayRequestBody
	payload, _ := json.Marshal(&request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "basic")
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

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
