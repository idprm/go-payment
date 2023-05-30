package nicepay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/utils"
)

type Nicepay struct {
	conf        *config.Secret
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
}

func NewNicepay(
	conf *config.Secret,
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
) *Nicepay {
	return &Nicepay{
		conf:        conf,
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
	}
}

func (p *Nicepay) Payment() ([]byte, error) {
	url := p.conf.Nicepay.Url + "/nicepay/direct/v2/registration"

	request := &entity.NicepayRequestBody{
		TimeStamp:         p.TimeStamp(),
		MerchantId:        p.conf.Nicepay.MerchantId,
		PaymentMethod:     "05",
		MitraCode:         p.order.Channel.Param,
		Currency:          p.gateway.GetCurrency(),
		PaymentAmount:     strconv.Itoa(int(p.order.GetAmount())),
		ReferenceNo:       p.order.GetNumber(),
		GoodsName:         p.order.GetDescription(),
		BuyerName:         p.order.GetName(),
		BuyerPhone:        p.order.GetMsisdn(),
		BuyerEmail:        p.order.GetEmail(),
		BuyerAddress:      "Billing Address",
		BuyerCity:         "Jakarta",
		BillingState:      "Jakarta",
		BillingPostNumber: "12345",
		BillingCountry:    "Indonesia",
		NotificationUrl:   p.conf.App.Url + "/v1/nicepay/notify",
		MerchantToken:     p.Token(),
	}

	if p.order.Channel.GetParam() == "OVOE" {
		request.CartData.Count = "1"
		request.CartData.NicepayRequestBodyItem = append(request.CartData.NicepayRequestBodyItem, "Consultation", "Consultation with Doctor", strconv.Itoa(int(p.order.GetAmount())), "1", "-")
	}

	payload, _ := json.Marshal(&request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
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

	p.logger.Writer(string(body))
	return body, nil
}

func (p *Nicepay) Token() string {
	valueToken := []byte(p.TimeStamp() + p.conf.Nicepay.MerchantId + p.order.Number + strconv.Itoa(int(p.order.Amount)) + p.conf.Nicepay.MerchantKey)
	return utils.EncryptSHA256(valueToken)
}

func (p *Nicepay) TimeStamp() string {
	return utils.TimeStamp()
}
