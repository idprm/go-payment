package midtrans

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
	APP_URL             string = utils.GetEnv("APP_URL")
	MIDTRANS_URL        string = utils.GetEnv("MIDTRANS_URL")
	MIDTRANS_MERCHANTID string = utils.GetEnv("MIDTRANS_MERCHANTID")
	MIDTRANS_CLIENTKEY  string = utils.GetEnv("MIDTRANS_CLIENTKEY")
	MIDTRANS_SERVERKEY  string = utils.GetEnv("MIDTRANS_SERVERKEY")
)

type Midtrans struct {
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
}

func NewMidtrans(
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
) *Midtrans {
	return &Midtrans{
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

/**
 * Payment Method
 */
func (p *Midtrans) Payment() ([]byte, error) {
	url := MIDTRANS_URL + "/transactions"

	var request entity.MidtransPaymentRequestBody
	request.ReqCustomer.FirstName = p.order.GetName()
	request.ReqCustomer.LastName = ""
	request.ReqCustomer.Phone = p.order.GetMsisdn()
	request.ReqCustomer.Email = p.order.GetEmail()
	request.ReqTransaction.OrderId = p.order.GetNumber()
	request.ReqTransaction.GrossAmount = int(p.order.Amount)
	request.ReqCallback.Finish = p.order.GetUrlReturn()

	payload, _ := json.Marshal(&request)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	var basic = "Basic " + base64.StdEncoding.EncodeToString([]byte(MIDTRANS_SERVERKEY))
	req.Header.Add("Authorization", basic)
	req.Header.Add("X-Override-Notification", APP_URL+"/v1/midtrans/notification")
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
func (p *Midtrans) Refund() ([]byte, error) {
	url := MIDTRANS_URL + "/refunds"

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	p.logger.Writer(string(body))
	return body, nil
}
