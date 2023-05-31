package midtrans

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

type Midtrans struct {
	conf        *config.Secret
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
}

func NewMidtrans(
	conf *config.Secret,
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
) *Midtrans {
	return &Midtrans{
		conf:        conf,
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
	}
}

type MidtransRequestBody struct {
	ReqTransaction struct {
		OrderId     string `json:"order_id"`
		GrossAmount int    `json:"gross_amount"`
	} `json:"transaction_details"`
	ReqCustomer struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
	} `json:"customer_details"`
	ReqCallback struct {
		Finish string `json:"finish"`
	} `json:"callbacks"`
}

// order
func (p *Midtrans) Payment() ([]byte, error) {
	url := p.conf.Midtrans.Url + "/transactions"

	var request entity.MidtransPaymentRequestBody
	request.ReqCustomer.FirstName = p.order.GetName()
	request.ReqCustomer.LastName = ""
	request.ReqCustomer.Phone = p.order.GetMsisdn()
	request.ReqCustomer.Email = p.order.GetEmail()
	request.ReqTransaction.OrderId = p.order.GetNumber()
	request.ReqTransaction.GrossAmount = int(p.order.Amount)
	request.ReqCallback.Finish = p.application.GetUrlReturn()

	payload, _ := json.Marshal(&request)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	var basic = "Basic " + base64.StdEncoding.EncodeToString([]byte(p.conf.Midtrans.ServerKey))
	req.Header.Add("Authorization", basic)
	req.Header.Add("X-Override-Notification", p.conf.App.Url+"/v1/midtrans/notification")
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

// refund
func (p *Midtrans) Refund() ([]byte, error) {
	url := p.conf.Midtrans.Url + "/refunds"

	var request entity.MidtransPaymentRequestBody
	request.ReqCustomer.FirstName = p.order.GetNumber()
	payload, _ := json.Marshal(&request)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

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
