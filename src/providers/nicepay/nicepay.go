package nicepay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/utils"
)

type Nicepay struct {
	conf  *config.Secret
	order *entity.Order
}

func NewNicepay(
	conf *config.Secret,
	order *entity.Order,
) *Nicepay {
	return &Nicepay{
		conf:  conf,
		order: order,
	}
}

type PaymentRequestBody struct {
	TimeStamp         string `json:"timeStamp"`
	MerchantId        string `json:"iMid"`
	PaymentMethod     string `json:"payMethod"`
	MitraCode         string `json:"mitraCd"`
	Currency          string `json:"currency"`
	PaymentAmount     string `json:"amt"`
	ReferenceNo       string `json:"referenceNo"`
	GoodsName         string `json:"goodsNm"`
	BuyerName         string `json:"billingNm"`
	BuyerPhone        string `json:"billingPhone"`
	BuyerEmail        string `json:"billingEmail"`
	BuyerAddress      string `json:"billingAddr"`
	BuyerCity         string `json:"billingCity"`
	BillingState      string `json:"billingState"`
	BillingPostNumber string `json:"billingPostCd"`
	BillingCountry    string `json:"billingCountry"`
	NotificationUrl   string `json:"dbProcessUrl"`
	MerchantToken     string `json:"merchantToken"`
	CartData          struct {
		Count string   `json:"count"`
		Item  []string `json:"item"`
	} `json:"cartData"`
}

type Item struct {
	GoodsName     string `json:"goods_name"`
	GoodsDetail   string `json:"goods_detail"`
	GoodsAmt      int    `json:"goods_amt"`
	GoodsQuantity int    `json:"goods_quantity"`
	ImgUrl        string `json:"img_url"`
}

func (p *Nicepay) Payment() ([]byte, error) {
	url := p.conf.Nicepay.Url + "/nicepay/direct/v2/registration"
	timeStamp := ""
	valueToken := []byte(timeStamp + p.conf.Nicepay.MerchantId + p.order.Number + strconv.Itoa(int(p.order.Amount)) + p.conf.Nicepay.MerchantKey)
	encryptToken := utils.EncryptSHA256(valueToken)

	request := &PaymentRequestBody{
		TimeStamp:         timeStamp,
		MerchantId:        p.conf.Nicepay.MerchantId,
		PaymentMethod:     "05",
		MitraCode:         p.order.Channel.Param,
		Currency:          "IDR",
		PaymentAmount:     strconv.Itoa(int(p.order.Amount)),
		ReferenceNo:       p.order.GetNumber(),
		GoodsName:         p.order.GetDescription(),
		BuyerName:         p.order.GetName(),
		BuyerPhone:        p.order.GetMsisdn(),
		BuyerEmail:        "help@sehatcepat.com",
		BuyerAddress:      "Billing Address",
		BuyerCity:         "Jakarta",
		BillingState:      "Jakarta",
		BillingPostNumber: "12345",
		BillingCountry:    "Indonesia",
		NotificationUrl:   p.conf.App.Url + "/v1/nicepay/notify",
		MerchantToken:     encryptToken,
	}

	// if order.Method.Param == "OVOE" {
	// 	request.CartData = "{}"
	// } else {
	// 	request.CartData = "{\"count\":\"1\",\"item\":[{\"goods_name\":\"Consultation\",\"goods_detail\":\"Consultation with Doctor\",\"goods_amt\":\"" + strconv.Itoa(order.Total) + "\",\"goods_quantity\":\"1\",\"img_url\":\"-\"}]}"
	// }

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
