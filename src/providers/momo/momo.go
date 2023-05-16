package momo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
)

type Momo struct {
	conf  *config.Secret
	order *entity.Order
}

func NewMomo(
	conf *config.Secret,
	order *entity.Order,
) *Momo {
	return &Momo{
		conf:  conf,
		order: order,
	}
}

type PaymentRequestBody struct {
	PartnerCode string `json:"partnerCode"`
	PartnerName string `json:"partnerName"`
	StoreId     string `json:"storeId"`
	RequestId   string `json:"requestId"`
	Amount      int    `json:"amount"`
	OrderId     string `json:"orderId"`
	OrderInfo   string `json:"orderInfo"`
	RedirectUrl string `json:"redirectUrl"`
	IpnUrl      string `json:"ipnUrl"`
	RequestType string `json:"requestType"`
	ExtraData   string `json:"extraData"`
	Lang        string `json:"lang"`
	AutoCapture bool   `json:"autoCapture"`
	Signature   string `json:"signature"`
}

type PaymentResponseBody struct {
	PartnerCode  string `json:"partnerCode"`
	OrderId      string `json:"orderId"`
	Amount       int    `json:"amount"`
	ResponseTime int    `json:"responseTime"`
	Message      string `json:"message"`
	ResultCode   int    `json:"resultCode"`
	PayUrl       string `json:"payUrl"`
}

type RefundRequestBody struct {
	PartnerCode string `json:"partnerCode"`
	OrderId     string `json:"orderId"`
	RequestId   string `json:"requestId"`
	Amount      int    `json:"amount"`
	TransId     int    `json:"transId"`
	Lang        string `json:"lang"`
	Description string `json:"description"`
	Signature   string `json:"signature"`
}

func (p *Momo) Payment() ([]byte, error) {
	url := p.conf.Momo.Url + "/v2/gateway/api/create"
	accessKey := p.conf.Momo.AccessKey
	partnerCode := p.conf.Momo.PartnerCode
	requestId := ""

	request := &PaymentRequestBody{
		PartnerName: "Test",
		PartnerCode: partnerCode,
		StoreId:     partnerCode,
		RequestId:   requestId,
		Amount:      int(p.order.Amount),
		OrderId:     p.order.Number,
		OrderInfo:   p.order.Description,
		RedirectUrl: "https://momo.vn",
		IpnUrl:      "https://webhook.site/94e534cb-a54a-4313-8e91-c42f7aa2e145",
		RequestType: "captureWallet",
		ExtraData:   "",
		Lang:        "en",
		AutoCapture: true,
		Signature:   p.HashTransaction(accessKey, int(p.order.Amount), "", "https://webhook.site/94e534cb-a54a-4313-8e91-c42f7aa2e145", p.order.Number, p.order.Description, partnerCode, "https://momo.vn", requestId, "captureWallet"),
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

func (p *Momo) Refund() ([]byte, error) {
	url := p.conf.Momo.Url + "/v2/gateway/api/refund"
	accessKey := p.conf.Momo.AccessKey
	partnerCode := p.conf.Momo.PartnerCode
	requestId := ""

	request := &RefundRequestBody{
		PartnerCode: partnerCode,
		OrderId:     p.order.Number,
		RequestId:   requestId,
		Amount:      int(p.order.Amount),
		TransId:     1683179398467,
		Lang:        "en",
		Description: p.order.Description,
	}

	request.Signature = p.HashRefund(accessKey, int(p.order.Amount), p.order.Description, p.order.Number, partnerCode, requestId, 1683179398467)

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

func (p *Momo) HashTransaction(accessKey string, amount int, extraData, ipnUrl, orderId, orderInfo, partnerCode, redirectUrl, requestId, requestType string) string {
	secret := p.conf.Momo.SecretKey
	data := "accessKey=" + accessKey + "&amount=" + strconv.Itoa(amount) + "&extraData=" + extraData + "&ipnUrl=" + ipnUrl + "&orderId=" + orderId + "&orderInfo=" + orderInfo + "&partnerCode=" + partnerCode + "&redirectUrl=" + redirectUrl + "&requestId=" + requestId + "&requestType=" + requestType
	// Create a new HMAC by defining the hash type and the key (as byte array)
	hm := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	hm.Write([]byte(data))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(hm.Sum(nil))
	return sha
}

func (p *Momo) HashRefund(accessKey string, amount int, description, orderId, partnerCode, requestId string, transId int) string {
	secret := p.conf.Momo.SecretKey
	data := "accessKey=" + accessKey + "&amount=" + strconv.Itoa(amount) + "&description=" + description + "&orderId=" + orderId + "&partnerCode=" + partnerCode + "&requestId=" + requestId + "&transId=" + strconv.Itoa(transId)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	hm := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	hm.Write([]byte(data))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(hm.Sum(nil))
	return sha
}
