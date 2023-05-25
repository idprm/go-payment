package jazzcash

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
)

type JazzCash struct {
	conf   *config.Secret
	logger *logger.Logger
	order  *entity.Order
}

func NewJazzCash(
	conf *config.Secret,
	logger *logger.Logger,
	order *entity.Order,
) *JazzCash {
	return &JazzCash{
		conf:   conf,
		logger: logger,
		order:  order,
	}
}

// payment
func (p *JazzCash) Payment() ([]byte, error) {
	merchantId := p.conf.JazzCash.MerchantId
	password := p.conf.JazzCash.Password

	orderInfo := strings.ReplaceAll(p.order.Number, "-", "")

	payload, _ := json.Marshal(&entity.JazzCashPaymentRequest{
		Language:          "EN",
		MerchantID:        p.conf.JazzCash.MerchantId,
		Password:          password,
		TxnRefNo:          p.order.GetNumber(),
		Amount:            strconv.Itoa(int(p.order.Amount)),
		TxnCurrency:       "PKR",
		TxnDateTime:       p.TxTime(),
		BillReference:     "billRef",
		Description:       p.order.GetDescription(),
		TxnExpiryDateTime: p.TxTimeExp(),
		SecureHash:        p.Hash(strconv.Itoa(int(p.order.Amount)), "billRef", 345678, p.order.GetDescription(), "EN", merchantId, p.PrefixMsisdn(), password, "PKR", orderInfo),
		MobileNumber:      p.PrefixMsisdn(),
		CNIC:              345678,
	})
	req, err := http.NewRequest("POST", p.conf.JazzCash.Url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, errors.New(err.Error())
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
		return nil, errors.New(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	p.logger.Writer(string(body))
	return body, nil
}

func (p *JazzCash) TxTime() string {
	currentTime := time.Now()
	return currentTime.Format("20060102150405")
}

func (p *JazzCash) TxTimeExp() string {
	tomorrowTime := time.Now().AddDate(0, 0, 1)
	return tomorrowTime.Format("20060102150405")
}

func (p *JazzCash) PrefixMsisdn() string {
	msisdn := p.order.Msisdn
	if strings.HasPrefix(msisdn, "92") {
		return "0" + strings.TrimPrefix(msisdn, "92")
	}
	return msisdn
}

func (p *JazzCash) Hash(amount string, billRef string, cnic int, description string, lang string, merchantId string, mobileNumber string, password string, currency string, orderInfo string) string {
	secret := p.conf.JazzCash.IntegeritySalt
	data := secret + "&" + amount + "&" + billRef + "&" + strconv.Itoa(cnic) + "&" + description + "&" + lang + "&" + merchantId + "&" + mobileNumber + "&" + password + "&" + currency + "&" + p.TxTime() + "&" + p.TxTimeExp() + "&" + orderInfo

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

func (p *JazzCash) SecureHashRefund(amount, merchantId, pin, password, currency, reffNo string) string {
	secret := p.conf.JazzCash.IntegeritySalt
	data := secret + "&" + amount + "&" + merchantId + "&" + pin + "&" + password + "&" + currency + "&" + reffNo

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
