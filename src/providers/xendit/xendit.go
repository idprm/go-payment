package xendit

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
	XENDIT_URL       string = utils.GetEnv("XENDIT_URL")
	XENDIT_SECRETKEY string = utils.GetEnv("XENDIT_SECRETKEY")
)

type Xendit struct {
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
}

func NewXendit(
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
) *Xendit {
	return &Xendit{
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
	}
}

func (p *Xendit) CreateInvoice() ([]byte, error) {
	url := XENDIT_URL + "/v2/invoices"
	request := &entity.XenditPayoutRequest{
		ExternalId: p.order.GetNumber(),
		Amount:     p.order.GetAmount(),
		Email:      "admin@sehatcepat.com",
	}
	request.SetInvoiceDuration(86400)
	request.SetSuccessRedirectUrl(p.order.GetUrlReturn())
	request.SetFailureRedirectUrl(p.order.GetUrlReturn())
	payload, _ := json.Marshal(&request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	var basic = "Basic " + base64.StdEncoding.EncodeToString([]byte(XENDIT_SECRETKEY))
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
