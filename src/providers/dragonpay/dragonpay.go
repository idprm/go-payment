package dragonpay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
)

type DragonPay struct {
	conf        *config.Secret
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
}

func NewDragonPay(
	conf *config.Secret,
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
) *DragonPay {
	return &DragonPay{
		conf:        conf,
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
	}
}

func (p *DragonPay) Payment() ([]byte, error) {
	url := p.conf.DragonPay.Url + p.order.GetNumber() + "/post"
	request := &entity.DragonPayRequestBody{
		Amount:      int(p.order.GetAmount()),
		Currency:    p.gateway.GetCurrency(),
		Description: p.order.GetDescription(),
		Email:       p.order.GetEmail(),
		MobileNo:    p.order.GetMsisdn(),
		ProcId:      p.channel.GetParam(),
		Param1:      p.application.GetUrlReturn(),
		Param2:      p.application.GetUrlReturn(),
		IpAddress:   p.order.GetIpAddress(),
	}
	payload, _ := json.Marshal(&request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	var basic = "Basic " + base64.StdEncoding.EncodeToString([]byte(p.conf.DragonPay.MerchantId+":"+p.conf.DragonPay.Password))
	req.Header.Add("Authorization", basic)
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

	p.logger.Writer(string(body))
	return body, nil
}

func (p *DragonPay) Refund() ([]byte, error) {
	return nil, nil
}
