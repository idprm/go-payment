package ximpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/utils"
	"github.com/idprm/go-payment/src/utils/hash_utils"
)

var (
	XIMPAY_URL_TSEL  string = utils.GetEnv("XIMPAY_URL_TSEL")
	XIMPAY_URL_H3I   string = utils.GetEnv("XIMPAY_URL_H3I")
	XIMPAY_URL_XL    string = utils.GetEnv("XIMPAY_URL_XL")
	XIMPAY_URL_ISAT  string = utils.GetEnv("XIMPAY_URL_ISAT")
	XIMPAY_URL_SF    string = utils.GetEnv("XIMPAY_URL_SF")
	XIMPAY_URL_XLPIN string = utils.GetEnv("XIMPAY_URL_XLPIN")
	XIMPAY_URL_SFPIN string = utils.GetEnv("XIMPAY_URL_SFPIN")
	XIMPAY_PARTNERID string = utils.GetEnv("XIMPAY_PARTNERID")
	XIMPAY_SECRETKEY string = utils.GetEnv("XIMPAY_SECRETKEY")
	XIMPAY_USERNAME  string = utils.GetEnv("XIMPAY_USERNAME")
)

type Ximpay struct {
	logger      *logger.Logger
	application *entity.Application
	gateway     *entity.Gateway
	channel     *entity.Channel
	order       *entity.Order
	payment     *entity.Payment
}

func NewXimpay(
	logger *logger.Logger,
	application *entity.Application,
	gateway *entity.Gateway,
	channel *entity.Channel,
	order *entity.Order,
	payment *entity.Payment,
) *Ximpay {
	return &Ximpay{
		logger:      logger,
		application: application,
		gateway:     gateway,
		channel:     channel,
		order:       order,
		payment:     payment,
	}
}

func (p *Ximpay) token(itemcode string) string {
	str := XIMPAY_PARTNERID + itemcode + p.order.GetNumber() + time.Now().Format("1/2/2006") + XIMPAY_SECRETKEY
	p.logger.Writer(strings.ToLower(str))
	return hash_utils.GetMD5Hash(strings.ToLower(str))
}

func (p *Ximpay) tokenSecond() string {
	tax := p.order.GetAmount() * 0.11
	str := XIMPAY_PARTNERID + fmt.Sprintf("%.0f", p.order.GetAmount()+tax) + p.order.GetNumber() + time.Now().Format("1/2/2006") + XIMPAY_SECRETKEY
	p.logger.Writer(strings.ToLower(str))
	return hash_utils.GetMD5Hash(strings.ToLower(str))
}

func (p *Ximpay) tokenWithoutTax() string {
	str := XIMPAY_PARTNERID + fmt.Sprintf("%.0f", p.order.GetAmount()) + p.order.GetNumber() + time.Now().Format("1/2/2006") + XIMPAY_SECRETKEY
	p.logger.Writer(strings.ToLower(str))
	return hash_utils.GetMD5Hash(strings.ToLower(str))
}

func (p *Ximpay) tokenPin(ximpayId, pin string) string {
	str := ximpayId + pin + XIMPAY_SECRETKEY
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
		url = XIMPAY_URL_TSEL
		req := &entity.XimpayTselRequestBody{
			PartnerId: XIMPAY_PARTNERID,
			CbParam:   p.order.GetNumber(),
			Op:        "TSEL",
			Msisdn:    p.order.GetMsisdn(),
		}

		if p.application.IsSuratSakit() || p.application.IsSurkit() {
			req.SetItemId(p.medicalCertificateHtiAndTselAmountToItemCode(p.order.GetAmount()))
			req.SetToken(p.token(p.medicalCertificateHtiAndTselAmountToItemCode(p.order.GetAmount())))
		}

		if p.application.IsSehatCepat() {
			req.SetItemId(p.consultationHtiAndTselAmountToItemCode(p.order.GetAmount()))
			req.SetToken(p.token(p.consultationHtiAndTselAmountToItemCode(p.order.GetAmount())))
		}

		payload, _ = json.Marshal(req)
	}

	if p.channel.IsHti() {
		url = XIMPAY_URL_H3I
		req := &entity.XimpayHtiRequestBody{
			PartnerId: XIMPAY_PARTNERID,
			CbParam:   p.order.GetNumber(),
			Op:        "HTI",
			Msisdn:    p.order.GetMsisdn(),
		}

		if p.application.IsSuratSakit() || p.application.IsSurkit() {
			req.SetItemId(p.medicalCertificateHtiAndTselAmountToItemCode(p.order.GetAmount()))
			req.SetToken(p.token(p.medicalCertificateHtiAndTselAmountToItemCode(p.order.GetAmount())))
		}

		if p.application.IsSehatCepat() {
			req.SetItemId(p.consultationHtiAndTselAmountToItemCode(p.order.GetAmount()))
			req.SetToken(p.token(p.consultationHtiAndTselAmountToItemCode(p.order.GetAmount())))
		}

		payload, _ = json.Marshal(req)
	}

	if p.channel.IsIsat() {
		url = XIMPAY_URL_ISAT
		// added tax 11%
		vat := int(p.order.GetAmount()) + int(p.order.GetAmount()*0.11)
		payload, _ = json.Marshal(
			&entity.XimpayIsatRequestBody{
				PartnerId:  XIMPAY_PARTNERID,
				ItemName:   "Item",
				ItemDesc:   "Item CEHAT",
				Amount:     vat,
				ChargeType: "ISAT_GENERAL",
				CbParam:    p.order.GetNumber(),
				Token:      p.tokenSecond(),
				Op:         "ISAT",
				Msisdn:     p.order.GetMsisdn(),
			},
		)
	}

	if p.channel.IsXl() {
		url = XIMPAY_URL_XL
		// added tax 11%
		vat := int(p.order.GetAmount()) + int(p.order.GetAmount()*0.11)
		payload, _ = json.Marshal(
			&entity.XimpayXlRequestBody{
				PartnerId: XIMPAY_PARTNERID,
				ItemName:  "Item",
				ItemDesc:  "Item CEHAT",
				Amount:    vat,
				CbParam:   p.order.GetNumber(),
				Token:     p.tokenSecond(),
				Op:        "xl",
				Msisdn:    p.order.GetMsisdn(),
			},
		)
	}

	if p.channel.IsSf() {
		url = XIMPAY_URL_SF
		payload, _ = json.Marshal(
			&entity.XimpaySfRequestBody{
				PartnerId: XIMPAY_PARTNERID,
				ItemName:  "Item",
				ItemDesc:  "Item CEHAT",
				AmountExc: int(p.order.GetAmount()),
				CbParam:   p.order.GetNumber(),
				Token:     p.tokenWithoutTax(),
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

func (p *Ximpay) Pin(ximpayId, pin string) ([]byte, error) {
	var url string
	var payload []byte

	if p.channel.IsXl() {
		url = XIMPAY_URL_XLPIN
		payload, _ = json.Marshal(
			&entity.XimpayPinRequestBody{
				XimpayId:    ximpayId,
				CodePin:     pin,
				XimpayToken: p.tokenPin(ximpayId, pin),
			},
		)
	}

	if p.channel.IsSf() {
		url = XIMPAY_URL_SFPIN
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

func (p *Ximpay) medicalCertificateHtiAndTselAmountToItemCode(amount float64) string {
	switch amount {
	case 5000:
		return "SHT01001"
	case 10000:
		return "SHT01002"
	case 15000:
		return "SHT01003"
	case 20000:
		return "SHT01004"
	case 25000:
		return "SHT01005"
	case 30000:
		return "SHT01006"
	case 35000:
		return "SHT01007"
	case 40000:
		return "SHT01008"
	case 45000:
		return "SHT01009"
	case 50000:
		return "SHT01010"
	case 55000:
		return "SHT01011"
	case 60000:
		return "SHT01012"
	case 65000:
		return "SHT01013"
	case 70000:
		return "SHT01014"
	case 75000:
		return "SHT01015"
	case 80000:
		return "SHT01016"
	case 85000:
		return "SHT01017"
	case 90000:
		return "SHT01018"
	case 95000:
		return "SHT01019"
	case 100000:
		return "SHT01020"
	}
	return "SHT01020"
}

func (p *Ximpay) consultationHtiAndTselAmountToItemCode(amount float64) string {
	switch amount {
	case 5000:
		return "SHT02001"
	case 10000:
		return "SHT02002"
	case 15000:
		return "SHT02003"
	case 20000:
		return "SHT02004"
	case 25000:
		return "SHT02005"
	case 30000:
		return "SHT02006"
	case 35000:
		return "SHT02007"
	case 40000:
		return "SHT02008"
	case 45000:
		return "SHT02009"
	case 50000:
		return "SHT02010"
	case 55000:
		return "SHT02011"
	case 60000:
		return "SHT02012"
	case 65000:
		return "SHT02013"
	case 70000:
		return "SHT02014"
	case 75000:
		return "SHT02015"
	case 80000:
		return "SHT02016"
	case 85000:
		return "SHT02017"
	case 90000:
		return "SHT02018"
	case 95000:
		return "SHT02019"
	case 100000:
		return "SHT02020"
	}
	return "SHT02020"
}
