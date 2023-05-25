package razer

import (
	"net/http"
	"strconv"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/utils"
)

type Razer struct {
	conf    *config.Secret
	logger  *logger.Logger
	order   *entity.Order
	channel *entity.Channel
}

func NewRazer(
	conf *config.Secret,
	logger *logger.Logger,
	order *entity.Order,
	channel *entity.Channel,
) *Razer {
	return &Razer{
		conf:    conf,
		logger:  logger,
		order:   order,
		channel: channel,
	}
}

func (p *Razer) Payment() (string, error) {
	url := p.conf.Razer.Url + "/RMS/pay/" + p.conf.Razer.MerchantId + "/" + p.channel.GetParam()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept-Charset", "utf-8")
	var str = strconv.Itoa(int(p.order.Amount)) + p.conf.Razer.MerchantId + p.order.Number + p.conf.Razer.VerifyKey
	q := req.URL.Query()
	q.Add("amount", strconv.Itoa(int(p.order.Amount)))
	q.Add("orderid", p.order.GetNumber())
	q.Add("bill_name", p.order.GetName())
	q.Add("bill_email", p.order.GetEmail())
	q.Add("bill_mobile", "+"+p.order.GetMsisdn())
	q.Add("bill_desc", p.order.GetDescription())
	q.Add("cur", "MYR")
	q.Add("vcode", utils.GetMD5Hash(str))
	p.logger.Writer(req)
	req.URL.RawQuery = q.Encode()
	returnUrl := url + "?" + q.Encode()
	p.logger.Writer(returnUrl)
	return returnUrl, nil
}
