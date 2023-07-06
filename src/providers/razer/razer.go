package razer

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
	"github.com/idprm/go-payment/src/utils/hash_utils"
)

type Razer struct {
	conf        *config.Secret
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
	payment     *entity.Payment
}

func NewRazer(
	conf *config.Secret,
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
	payment *entity.Payment,
) *Razer {
	return &Razer{
		conf:        conf,
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
		payment:     payment,
	}
}

/**
 * Payment Method
 */
func (p *Razer) Payment() (string, error) {
	url := p.conf.Razer.Url + "/RMS/pay/" + p.conf.Razer.MerchantId + "/" + p.channel.GetParam()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		p.logger.Writer(err)
		return "", err
	}
	req.Header.Add("Accept-Charset", "utf-8")
	var str = strconv.Itoa(int(p.order.Amount)) + p.conf.Razer.MerchantId + p.order.Number + p.conf.Razer.VerifyKey
	q := req.URL.Query()
	q.Add("amount", strconv.Itoa(int(p.order.Amount)))
	q.Add("orderid", p.order.GetNumber())
	q.Add("bill_name", p.order.GetName())
	q.Add("bill_email", p.order.GetEmail())
	q.Add("bill_mobile", "+"+p.order.GetMsisdn())
	q.Add("bill_desc", p.order.GetDescription())
	q.Add("cur", p.gateway.GetCurrency())
	q.Add("vcode", hash_utils.GetMD5Hash(str))
	p.logger.Writer(req)
	req.URL.RawQuery = q.Encode()
	returnUrl := url + "?" + q.Encode()
	p.logger.Writer(returnUrl)
	return returnUrl, nil
}

/**
 * Refund Method
 */
func (p *Razer) Refund() ([]byte, error) {
	url := p.conf.Razer.UrlApi + "/RMS/API/refundAPI/refund.php"
	request := &entity.RefundRazerRequestBody{
		TransactionId: p.payment.GetTransactionId(),
		MerchantID:    p.conf.Razer.MerchantId,
	}
	var str = p.payment.GetTransactionId() + p.conf.Razer.MerchantId + p.conf.Razer.VerifyKey
	request.SetSignature(hash_utils.GetMD5Hash(str))
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
