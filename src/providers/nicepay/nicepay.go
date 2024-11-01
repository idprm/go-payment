package nicepay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/utils"
	"github.com/idprm/go-payment/src/utils/hash_utils"
)

var (
	APP_URL             string = utils.GetEnv("APP_URL")
	NICEPAY_URL         string = utils.GetEnv("NICEPAY_URL")
	NICEPAY_MERCHANTID  string = utils.GetEnv("NICEPAY_MERCHANTID")
	NICEPAY_MERCHANTKEY string = utils.GetEnv("NICEPAY_MERCHANTKEY")
)

type Nicepay struct {
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
}

func NewNicepay(
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
) *Nicepay {
	return &Nicepay{
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
	}
}

/**
 * Payment Method
 */
func (p *Nicepay) Payment() ([]byte, error) {
	url := NICEPAY_URL + "/nicepay/direct/v2/registration"

	request := &entity.NicepayRequestBody{
		TimeStamp:         p.TimeStamp(),
		MerchantId:        NICEPAY_MERCHANTID,
		PaymentMethod:     "05",
		MitraCode:         p.channel.GetParam(),
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
		NotificationUrl:   APP_URL + "/v1/nicepay/notification",
		MerchantToken:     p.Token(),
	}

	if p.channel.GetParam() == "OVOE" {
		request.CartData = "{}"
	} else {
		request.CartData = "{\"count\":\"1\",\"item\":[{\"goods_name\":\"Consultation\",\"goods_detail\":\"Consultation with Doctor\",\"goods_amt\":\"" + strconv.Itoa(int(p.order.Amount)) + "\",\"goods_quantity\":\"1\",\"img_url\":\"-\"}]}"
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	p.logger.Writer(string(body))
	return body, nil
}

/**
 * Refund Method
 */
func (p *Nicepay) Refund() ([]byte, error) {
	return nil, nil
}

func (p *Nicepay) Token() string {
	valueToken := []byte(p.TimeStamp() + NICEPAY_MERCHANTID + p.order.Number + strconv.Itoa(int(p.order.Amount)) + NICEPAY_MERCHANTKEY)
	return hash_utils.EncryptSHA256(valueToken)
}

func (p *Nicepay) TimeStamp() string {
	return hash_utils.TimeStamp()
}
