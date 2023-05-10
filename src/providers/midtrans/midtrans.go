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
)

type Midtrans struct {
	conf  *config.Secret
	order *entity.Order
}

func NewMidtrans(
	conf *config.Secret,
	order *entity.Order,
) *Midtrans {
	return &Midtrans{
		conf:  conf,
		order: order,
	}
}

// payment
func (p *Midtrans) Payment() ([]byte, error) {
	url := p.conf.Midtrans.Url + "/transactions"

	var request entity.MidtransPaymentRequestBody
	request.ReqCustomer.FirstName = p.order.GetNumber()
	request.ReqCustomer.LastName = ""
	request.ReqCustomer.Phone = p.order.GetMsisdn()
	request.ReqCustomer.Email = p.order.GetEmail()
	request.ReqTransaction.OrderId = p.order.GetNumber()
	request.ReqTransaction.GrossAmount = int(p.order.Amount)
	request.ReqCallback.Finish = ""

	payload, _ := json.Marshal(&request)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(p.conf.Midtrans.ServerKey)))
	req.Header.Add("X-Override-Notification", p.conf.App.Url+"/midtrans/notification")
	req.Header.Set("Content-Type", "application/json")
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
