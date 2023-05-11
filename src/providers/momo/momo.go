package momo

import (
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

// func (p *Momo) Payment() {
// 	url := config.ViperEnv("MOMO_URL") + "/v2/gateway/api/create"
// 	accessKey := config.ViperEnv("MOMO_ACCESS_KEY")
// 	partnerCode := config.ViperEnv("MOMO_PARTNER_CODE")
// 	requestId := helper.GenerateTransactionId()

// 	var request dto.MomoTransactionRequestBody
// 	request.PartnerName = "Test"
// 	request.PartnerCode = partnerCode
// 	request.StoreId = partnerCode
// 	request.RequestId = requestId
// 	request.Amount = h.order.Total
// 	request.OrderId = h.order.Number
// 	request.OrderInfo = "Games"
// 	request.RedirectUrl = "https://momo.vn"
// 	request.IpnUrl = "https://webhook.site/94e534cb-a54a-4313-8e91-c42f7aa2e145"
// 	request.RequestType = "captureWallet"
// 	request.ExtraData = ""
// 	request.Lang = "en"
// 	request.AutoCapture = true
// 	request.Signature = h.HashTransaction(accessKey, h.order.Total, "", "https://webhook.site/94e534cb-a54a-4313-8e91-c42f7aa2e145", h.order.Number, "Games", partnerCode, "https://momo.vn", requestId, "captureWallet")

// 	payload, _ := json.Marshal(&request)
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
// 	if err != nil {
// 		return nil, errors.New(err.Error())
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	log.Println(req)

// 	helper.Log.WithFields(logrus.Fields{"payload": datatypes.JSON(payload)}).Info("REQUEST BODY MOMO")

// 	tr := &http.Transport{
// 		Proxy:              http.ProxyFromEnvironment,
// 		MaxIdleConns:       10,
// 		IdleConnTimeout:    30 * time.Second,
// 		DisableCompression: true,
// 	}

// 	client := &http.Client{
// 		Timeout:   30 * time.Second,
// 		Transport: tr,
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, errors.New(err.Error())
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, errors.New(err.Error())
// 	}

// 	return body, nil
// }
