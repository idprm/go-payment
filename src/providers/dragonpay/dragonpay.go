package dragonpay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/utils"
)

var (
	DRAGONPAY_URL        string = utils.GetEnv("DRAGONPAY_URL")
	DRAGONPAY_MERCHANTID string = utils.GetEnv("DRAGONPAY_MERCHANTID")
	DRAGONPAY_PASSWORD   string = utils.GetEnv("DRAGONPAY_PASSWORD")
)

type DragonPay struct {
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
}

func NewDragonPay(
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
) *DragonPay {
	return &DragonPay{
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
	}
}

func (p *DragonPay) Payment() ([]byte, error) {
	url := DRAGONPAY_URL + p.order.GetNumber() + "/post"
	request := &entity.DragonPayRequestBody{
		Amount:      int(p.order.GetAmount()),
		Currency:    p.gateway.GetCurrency(),
		Description: p.order.GetDescription(),
		Email:       p.order.GetEmail(),
		MobileNo:    p.order.GetMsisdn(),
		ProcId:      p.channel.GetParam(),
		Param1:      p.order.GetUrlReturn(),
		Param2:      p.order.GetUrlReturn(),
		IpAddress:   p.order.GetIpAddress(),
	}
	payload, _ := json.Marshal(&request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	var basic = "Basic " + base64.StdEncoding.EncodeToString([]byte(DRAGONPAY_MERCHANTID+":"+DRAGONPAY_PASSWORD))
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	p.logger.Writer(string(body))
	return body, nil
}

func (p *DragonPay) Refund() ([]byte, error) {
	return nil, nil
}
