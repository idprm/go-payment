package ximpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/utils/hash_utils"
	"github.com/sirupsen/logrus"
)

type Ximpay struct {
	conf        *config.Secret
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
	payment     *entity.Payment
}

func NewXimpay(
	conf *config.Secret,
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
	payment *entity.Payment,
) *Ximpay {
	return &Ximpay{
		conf:        conf,
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
		payment:     payment,
	}
}

func (p *Ximpay) token() string {
	str := p.conf.Ximpay.PartnerId + "SHT00001" + p.order.GetNumber() + time.Now().Format("1/2/2006") + p.conf.Ximpay.SecretKey
	l := p.logger.Init("order", true)
	l.WithFields(logrus.Fields{"plain_text": str}).Info("TOKEN")
	return hash_utils.GetMD5Hash(strings.ToLower(str))
}

func (p *Ximpay) tokenSecond() string {
	str := p.conf.Ximpay.PartnerId + fmt.Sprintf("%f", p.order.GetAmount()) + p.order.GetNumber() + time.Now().Format("1/2/2006") + p.conf.Ximpay.SecretKey
	return hash_utils.GetMD5Hash(strings.ToLower(str))
}

/**
 * Payment Method
 */
func (p *Ximpay) Payment() ([]byte, error) {
	var url string
	var payload []byte

	if p.channel.IsTsel() {
		url = p.conf.Ximpay.UrlTsel
		payload, _ = json.Marshal(
			&entity.XimpayTselRequestBody{
				PartnerId: p.conf.Ximpay.PartnerId,
				ItemId:    "SHT00001",
				CbParam:   p.order.GetNumber(),
				Token:     p.token(),
				Op:        "TSEL",
				Msisdn:    p.order.GetMsisdn(),
			},
		)
	}

	if p.channel.IsHti() {
		url = p.conf.Ximpay.UrlHti
		payload, _ = json.Marshal(
			&entity.XimpayHtiRequestBody{
				PartnerId: p.conf.Ximpay.PartnerId,
				ItemId:    "SHT00001",
				CbParam:   p.order.GetNumber(),
				Token:     p.token(),
				Op:        "HTI",
				Msisdn:    p.order.GetMsisdn(),
			},
		)
	}

	if p.channel.IsIsat() {
		url = p.conf.Ximpay.UrlIsat
		payload, _ = json.Marshal(
			&entity.XimpayIsatRequestBody{
				PartnerId:  p.conf.Ximpay.PartnerId,
				ItemName:   "Item 2K",
				ItemDesc:   "Item 2K",
				Amount:     int(p.order.GetAmount()),
				ChargeType: "ISAT_GENERAL",
				CbParam:    p.order.GetNumber(),
				Token:      p.token(),
				Op:         "ISAT",
				Msisdn:     p.order.GetMsisdn(),
			},
		)
	}

	if p.channel.IsXl() {
		url = p.conf.Ximpay.UrlXl
		payload, _ = json.Marshal(
			&entity.XimpayXlRequestBody{
				PartnerId: p.conf.Ximpay.PartnerId,
				ItemName:  "Item 2K",
				ItemDesc:  "Item 2K",
				Amount:    int(p.order.GetAmount()),
				CbParam:   p.order.GetNumber(),
				Token:     p.tokenSecond(),
				Op:        "xl",
				Msisdn:    p.order.GetMsisdn(),
			},
		)
	}

	if p.channel.IsSf() {
		url = p.conf.Ximpay.UrlSf
		payload, _ = json.Marshal(
			&entity.XimpayXlRequestBody{
				PartnerId: p.conf.Ximpay.PartnerId,
				ItemName:  "Item 2K",
				ItemDesc:  "Item 2K",
				Amount:    int(p.order.GetAmount()),
				CbParam:   p.order.GetNumber(),
				Token:     p.tokenSecond(),
				Op:        "SF",
				Msisdn:    p.order.GetMsisdn(),
			},
		)
	}

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

func (p *Ximpay) Pin() ([]byte, error) {
	var url string
	var payload []byte

	if p.channel.IsXl() {
		url = p.conf.Ximpay.UrlXlPin
		payload, _ = json.Marshal(
			&entity.XimpayPinRequestBody{
				XimpayId:    "",
				CodePin:     "",
				XimpayToken: "",
			},
		)
	}

	if p.channel.IsSf() {
		url = p.conf.Ximpay.UrlSfPin
		payload, _ = json.Marshal(
			&entity.XimpayPinRequestBody{
				XimpayId:    "",
				CodePin:     "",
				XimpayToken: "",
			},
		)
	}

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
