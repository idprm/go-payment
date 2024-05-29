package momo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
	APP_URL          string = utils.GetEnv("APP_URL")
	MOMO_URL         string = utils.GetEnv("MOMO_URL")
	MOMO_PARTNERCODE string = utils.GetEnv("MOMO_PARTNERCODE")
	MOMO_ACCESSKEY   string = utils.GetEnv("MOMO_ACCESSKEY")
	MOMO_SECRETKEY   string = utils.GetEnv("MOMO_SECRETKEY")
)

type Momo struct {
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
}

func NewMomo(
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
) *Momo {
	return &Momo{
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
func (p *Momo) Payment() ([]byte, error) {
	url := MOMO_URL + "/v2/gateway/api/create"
	accessKey := MOMO_ACCESSKEY
	partnerCode := MOMO_PARTNERCODE
	requestId := hash_utils.GenerateTransactionId()

	request := &entity.MomoRequestBody{
		PartnerName: "Test",
		PartnerCode: partnerCode,
		StoreId:     partnerCode,
		RequestId:   requestId,
		Amount:      int(p.order.GetAmount()),
		OrderId:     p.order.GetNumber(),
		OrderInfo:   p.order.GetDescription(),
		RedirectUrl: p.order.GetUrlReturn(),
		IpnUrl:      APP_URL + "/v1/momo/notification",
		RequestType: "captureWallet",
		ExtraData:   "",
		Lang:        "en",
		AutoCapture: true,
		Signature:   p.HashTransaction(accessKey, int(p.order.Amount), "", APP_URL+"/v1/momo/notification", p.order.GetNumber(), p.order.GetDescription(), partnerCode, p.order.GetUrlReturn(), requestId, "captureWallet"),
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
func (p *Momo) Refund() ([]byte, error) {
	url := MOMO_URL + "/v2/gateway/api/refund"
	accessKey := MOMO_ACCESSKEY
	partnerCode := MOMO_PARTNERCODE
	requestId := hash_utils.GenerateTransactionId()

	request := &entity.MomoRefundRequestBody{
		PartnerCode: partnerCode,
		OrderId:     p.order.GetNumber(),
		RequestId:   requestId,
		Amount:      int(p.order.GetAmount()),
		TransId:     1683179398467,
		Lang:        p.gateway.Country.GetLocale(),
		Description: p.order.GetDescription(),
	}

	request.Signature = p.HashRefund(accessKey, int(p.order.Amount), p.order.GetDescription(), p.order.GetNumber(), partnerCode, requestId, 1683179398467)

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

func (p *Momo) HashTransaction(accessKey string, amount int, extraData, ipnUrl, orderId, orderInfo, partnerCode, redirectUrl, requestId, requestType string) string {
	secret := MOMO_ACCESSKEY
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
	secret := MOMO_ACCESSKEY
	data := "accessKey=" + accessKey + "&amount=" + strconv.Itoa(amount) + "&description=" + description + "&orderId=" + orderId + "&partnerCode=" + partnerCode + "&requestId=" + requestId + "&transId=" + strconv.Itoa(transId)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	hm := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	hm.Write([]byte(data))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(hm.Sum(nil))
	return sha
}
