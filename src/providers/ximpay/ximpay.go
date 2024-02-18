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
	"github.com/idprm/go-payment/src/services"
	"github.com/idprm/go-payment/src/utils/hash_utils"
)

type Ximpay struct {
	conf          *config.Secret
	logger        *logger.Logger
	application   *entity.Application
	gateway       *entity.Gateway
	channel       *entity.Channel
	order         *entity.Order
	payment       *entity.Payment
	verifyService services.IVerifyService
}

func NewXimpay(
	conf *config.Secret,
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
	payment *entity.Payment,
	verifyService services.IVerifyService,
) *Ximpay {
	return &Ximpay{
		conf:          conf,
		logger:        logger,
		application:   application,
		gateway:       gateway,
		channel:       channel,
		order:         order,
		payment:       payment,
		verifyService: verifyService,
	}
}

func (p *Ximpay) token() string {
	str := p.conf.Ximpay.PartnerId + "SHT00001" + p.order.GetNumber() + time.Now().Format("1/2/2006") + p.conf.Ximpay.SecretKey
	p.logger.Writer(strings.ToLower(str))
	return hash_utils.GetMD5Hash(strings.ToLower(str))
}

func (p *Ximpay) tokenSecond() string {
	str := p.conf.Ximpay.PartnerId + fmt.Sprintf("%.0f", p.order.GetAmount()) + p.order.GetNumber() + time.Now().Format("1/2/2006") + p.conf.Ximpay.SecretKey
	p.logger.Writer(strings.ToLower(str))
	return hash_utils.GetMD5Hash(strings.ToLower(str))
}

func (p *Ximpay) tokenPin(ximpayId, pin string) string {
	str := ximpayId + pin + p.conf.Ximpay.SecretKey
	p.logger.Writer(strings.ToLower(str))
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
		p.verifyService.Set(&entity.Verify{
			Key:  p.order.GetNumber(),
			Data: p.token(),
		})
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
		p.verifyService.Set(&entity.Verify{
			Key:  p.order.GetNumber(),
			Data: p.token(),
		})
	}

	if p.channel.IsIsat() {
		url = p.conf.Ximpay.UrlIsat
		payload, _ = json.Marshal(
			&entity.XimpayIsatRequestBody{
				PartnerId:  p.conf.Ximpay.PartnerId,
				ItemName:   "Item 2K",
				ItemDesc:   "Item 2K CEHAT",
				Amount:     int(p.order.GetAmount()),
				ChargeType: "ISAT_GENERAL",
				CbParam:    p.order.GetNumber(),
				Token:      p.tokenSecond(),
				Op:         "ISAT",
				Msisdn:     p.order.GetMsisdn(),
			},
		)
		p.verifyService.Set(&entity.Verify{
			Key:  p.order.GetNumber(),
			Data: p.token(),
		})
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
		p.verifyService.Set(&entity.Verify{
			Key:  p.order.GetNumber(),
			Data: p.token(),
		})
	}

	if p.channel.IsSf() {
		url = p.conf.Ximpay.UrlSf
		payload, _ = json.Marshal(
			&entity.XimpaySfRequestBody{
				PartnerId: p.conf.Ximpay.PartnerId,
				ItemName:  "Item 2K",
				ItemDesc:  "Item 2K",
				AmountExc: int(p.order.GetAmount()),
				CbParam:   p.order.GetNumber(),
				Token:     p.tokenSecond(),
				Op:        "SF",
				Msisdn:    p.order.GetMsisdn(),
			},
		)
		p.verifyService.Set(&entity.Verify{
			Key:  p.order.GetNumber(),
			Data: p.token(),
		})
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

func (p *Ximpay) Pin(ximpayId, pin string) ([]byte, error) {
	var url string
	var payload []byte

	if p.channel.IsXl() {
		url = p.conf.Ximpay.UrlXlPin
		payload, _ = json.Marshal(
			&entity.XimpayPinRequestBody{
				XimpayId:    ximpayId,
				CodePin:     pin,
				XimpayToken: p.tokenPin(ximpayId, pin),
			},
		)
	}

	if p.channel.IsSf() {
		url = p.conf.Ximpay.UrlSfPin
		payload, _ = json.Marshal(
			&entity.XimpayPinRequestBody{
				XimpayId:    ximpayId,
				CodePin:     pin,
				XimpayToken: p.tokenPin(ximpayId, pin),
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
