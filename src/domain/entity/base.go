package entity

import (
	"net/http"
	"strings"
	"time"

	"github.com/idprm/go-payment/src/utils/hash_utils"
)

type ErrorResponse struct {
	Field string
	Tag   string
	Value string
}

type (
	OrderBodyRequest struct {
		UrlCallback string  `validate:"required" json:"url_callback"`
		UrlReturn   string  `validate:"required" json:"url_return"`
		Msisdn      string  `validate:"required" json:"msisdn"`
		Name        string  `validate:"required" json:"name"`
		Number      string  `validate:"required" json:"number"`
		Channel     string  `validate:"required" json:"channel"`
		Amount      float64 `validate:"required" json:"amount"`
		Email       string  `json:"email"`
		Description string  `json:"description"`
		IpAddress   string  `json:"ip_address"`
	}
	OrderBodyResponse struct {
		Error       bool   `json:"error"`
		StatusCode  int    `json:"status_code"`
		Message     string `json:"message"`
		RedirectUrl string `json:"redirect_url"`
	}

	OrderBodyWithoutRedirectResponse struct {
		Error      bool   `json:"error"`
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}

	OrderPINBodyRequest struct {
		UrlCallback string `validate:"required" json:"url_callback"`
		UrlReturn   string `validate:"required" json:"url_return"`
		Channel     string `validate:"required" json:"channel"`
		Msisdn      string `validate:"required" json:"msisdn"`
		PIN         string `validate:"required" json:"pin"`
	}
)

type PaymentBodyResponse struct {
	Error      bool   `json:"error"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewStatusOKPaymentBodyResponse() *PaymentBodyResponse {
	return &PaymentBodyResponse{
		Error:      false,
		StatusCode: http.StatusOK,
		Message:    "success",
	}
}

func NewStatusCreatedPaymentBodyResponse() *PaymentBodyResponse {
	return &PaymentBodyResponse{
		Error:      false,
		StatusCode: http.StatusCreated,
		Message:    "success",
	}
}

func NewStatusCreatedOrderBodyResponse(url string) *OrderBodyResponse {
	return &OrderBodyResponse{
		Error:       false,
		StatusCode:  http.StatusCreated,
		Message:     "success",
		RedirectUrl: url,
	}
}

func NewStatusCreatedOrderBodyMessageResponse(message string) *OrderBodyWithoutRedirectResponse {
	return &OrderBodyWithoutRedirectResponse{
		Error:      false,
		StatusCode: http.StatusCreated,
		Message:    message,
	}
}

func (r *OrderBodyRequest) GetUrlCallback() string {
	return r.UrlCallback
}

func (r *OrderBodyRequest) GetUrlReturn() string {
	return r.UrlReturn
}

func (r *OrderBodyRequest) GetChannel() string {
	return r.Channel
}

func (r *OrderBodyRequest) GetMsisdn() string {
	return r.Msisdn
}

func (r *OrderBodyRequest) GetName() string {
	return r.Name
}

func (r *OrderBodyRequest) GetEmail() string {
	return r.Email
}

func (r *OrderBodyRequest) GetNumber() string {
	return r.Number
}

func (r *OrderBodyRequest) GetAmount() float64 {
	return r.Amount
}

func (r *OrderBodyRequest) GetDescription() string {
	return r.Description
}

func (r *OrderBodyRequest) GetIpAddress() string {
	return r.IpAddress
}

func (r *OrderPINBodyRequest) GetUrlCallback() string {
	return r.UrlCallback
}

func (r *OrderPINBodyRequest) GetUrlReturn() string {
	return r.UrlReturn
}

func (r *OrderPINBodyRequest) GetChannel() string {
	return r.Channel
}

func (r *OrderPINBodyRequest) GetMsisdn() string {
	return r.Msisdn
}

func (r *OrderPINBodyRequest) GetPIN() string {
	return r.PIN
}

type RefundRequestBody struct {
	Number      string `json:"number"`
	UrlCallback string `json:"url_callback"`
}

func (r *RefundRequestBody) GetNumber() string {
	return r.Number
}

func (r *RefundRequestBody) GetUrlCallback() string {
	return r.UrlCallback
}

type DragonPayRequestBody struct {
	Amount      int    `json:"Amount"`
	Currency    string `json:"Currency"`
	Description string `json:"Description"`
	Email       string `json:"Email"`
	MobileNo    string `json:"MobileNo"`
	ProcId      string `json:"ProcId"`
	Param1      string `json:"Param1"`
	Param2      string `json:"Param2"`
	IpAddress   string `json:"IpAddress"`
}

func (r *DragonPayRequestBody) GetAmount() int {
	return r.Amount
}

func (r *DragonPayRequestBody) GetCurrency() string {
	return r.Currency
}

func (r *DragonPayRequestBody) GetDescription() string {
	return r.Description
}

func (r *DragonPayRequestBody) GetEmail() string {
	return r.Email
}

func (r *DragonPayRequestBody) GetMobileNo() string {
	return r.MobileNo
}

func (r *DragonPayRequestBody) GetProcId() string {
	return r.ProcId
}

func (r *DragonPayRequestBody) GetParam1() string {
	return r.Param1
}

func (r *DragonPayRequestBody) GetParam2() string {
	return r.Param2
}

func (r *DragonPayRequestBody) GetIpAddress() string {
	return r.IpAddress
}

type DragonPayResponsePayload struct {
	RefNo   string `json:"RefNo"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
	Url     string `json:"Url"`
}

func (r *DragonPayResponsePayload) GetRefNo() string {
	return r.RefNo
}

func (r *DragonPayResponsePayload) GetStatus() string {
	return r.Status
}

func (r *DragonPayResponsePayload) GetMessage() string {
	return r.Message
}

func (r *DragonPayResponsePayload) GetUrl() string {
	return r.Url
}

func (r *DragonPayResponsePayload) IsValid() bool {
	return r.GetUrl() != ""
}

type NotifDragonPayRequestBody struct {
	TransactionId string  `json:"txnid" xml:"txnid" form:"txnid"`
	ReferenceNo   string  `json:"refno" xml:"refno" form:"refno"`
	Status        string  `json:"status" xml:"status" form:"status"`
	Message       string  `json:"message" xml:"message" form:"message"`
	Amount        float64 `json:"amount" xml:"amount" form:"amount"`
	Currency      string  `json:"ccy" xml:"ccy" form:"ccy"`
	ProcId        string  `json:"procid" xml:"procid" form:"procid"`
	Digest        string  `json:"digest" xml:"digest" form:"digest"`
}

func (r *NotifDragonPayRequestBody) GetTransactionId() string {
	return r.TransactionId
}

func (r *NotifDragonPayRequestBody) GetReferenceNo() string {
	return r.ReferenceNo
}

func (r *NotifDragonPayRequestBody) GetStatus() string {
	return r.Status
}

func (r *NotifDragonPayRequestBody) GetMessage() string {
	return r.Message
}

func (r *NotifDragonPayRequestBody) GetAmount() float64 {
	return r.Amount
}

func (r *NotifDragonPayRequestBody) GetCurrency() string {
	return r.Currency
}

func (r *NotifDragonPayRequestBody) GetProcId() string {
	return r.ProcId
}

func (r *NotifDragonPayRequestBody) GetDigest() string {
	return r.Digest
}

func (r *NotifDragonPayRequestBody) IsValid() bool {
	return r.GetReferenceNo() != "" || r.GetStatus() != ""
}

type JazzCashPaymentRequest struct {
	Language          string `json:"pp_Language"`
	MerchantID        string `json:"pp_MerchantID"`
	SubMerchantID     string `json:"pp_SubMerchantID"`
	Password          string `json:"pp_Password"`
	TxnRefNo          string `json:"pp_TxnRefNo"`
	Amount            string `json:"pp_Amount"`
	TxnCurrency       string `json:"pp_TxnCurrency"`
	TxnDateTime       string `json:"pp_TxnDateTime"`
	BillReference     string `json:"pp_BillReference"`
	Description       string `json:"pp_Description"`
	TxnExpiryDateTime string `json:"pp_TxnExpiryDateTime"`
	MobileNumber      string `json:"pp_MobileNumber"`
	SecureHash        string `json:"pp_SecureHash"`
	CNIC              int    `json:"pp_CNIC"`
}

func (r *JazzCashPaymentRequest) GetLanguage() string {
	return r.Language
}

func (r *JazzCashPaymentRequest) GetMerchantID() string {
	return r.MerchantID
}

func (r *JazzCashPaymentRequest) GetSubMerchantID() string {
	return r.SubMerchantID
}

func (r *JazzCashPaymentRequest) GetPassword() string {
	return r.Password
}

func (r *JazzCashPaymentRequest) GetTxnRefNo() string {
	return r.TxnRefNo
}

func (r *JazzCashPaymentRequest) GetAmount() string {
	return r.Amount
}

func (r *JazzCashPaymentRequest) GetTxnCurrency() string {
	return r.TxnCurrency
}

func (r *JazzCashPaymentRequest) GetTxnDateTime() string {
	return r.TxnDateTime
}

func (r *JazzCashPaymentRequest) GetBillReference() string {
	return r.BillReference
}

func (r *JazzCashPaymentRequest) GetDescription() string {
	return r.Description
}

func (r *JazzCashPaymentRequest) GetTxnExpiryDateTime() string {
	return r.TxnExpiryDateTime
}

func (r *JazzCashPaymentRequest) GetMobileNumber() string {
	return r.MobileNumber
}

func (r *JazzCashPaymentRequest) GetSecureHash() string {
	return r.SecureHash
}

func (r *JazzCashPaymentRequest) GetCNIC() int {
	return r.CNIC
}

type JazzCashRefundRequest struct {
	TxnRefNo     string `json:"pp_TxnRefNo"`
	Amount       string `json:"pp_Amount"`
	TxnCurrency  string `json:"pp_TxnCurrency"`
	MerchantID   string `json:"pp_MerchantID"`
	Password     string `json:"pp_Password"`
	MerchantMPIN string `json:"pp_MerchantMPIN"`
	SecureHash   string `json:"pp_SecureHash"`
}

func (r *JazzCashRefundRequest) GetTxnRefNo() string {
	return r.TxnRefNo
}

func (r *JazzCashRefundRequest) GetAmount() string {
	return r.Amount
}

func (r *JazzCashRefundRequest) GetTxnCurrency() string {
	return r.TxnCurrency
}

func (r *JazzCashRefundRequest) GetMerchantID() string {
	return r.MerchantID
}

func (r *JazzCashRefundRequest) GetPassword() string {
	return r.Password
}

func (r *JazzCashRefundRequest) GetMerchantMPIN() string {
	return r.MerchantMPIN
}

func (r *JazzCashRefundRequest) GetSecureHash() string {
	return r.SecureHash
}

type JazzCashRefundResponse struct {
	SecureHash      string `json:"pp_SecureHash"`
	ResponseCode    string `json:"pp_ResponseCode"`
	ResponseMessage string `json:"pp_ResponseMessage"`
}

func (r *JazzCashRefundResponse) GetSecureHash() string {
	return r.SecureHash
}

func (r *JazzCashRefundResponse) GetResponseCode() string {
	return r.ResponseCode
}

func (r *JazzCashRefundResponse) GetResponseMessage() string {
	return r.ResponseMessage
}

type JazzCashInquiryRequest struct {
	TxnRefNo   string `json:"pp_TxnRefNo"`
	MerchantID string `json:"pp_MerchantID"`
	Password   string `json:"pp_Password"`
	SecureHash string `json:"pp_SecureHash"`
}

func (r *JazzCashInquiryRequest) GetTxnRefNo() string {
	return r.TxnRefNo
}

func (r *JazzCashInquiryRequest) GetMerchantID() string {
	return r.MerchantID
}

func (r *JazzCashInquiryRequest) GetPassword() string {
	return r.Password
}

func (r *JazzCashInquiryRequest) GetSecureHash() string {
	return r.SecureHash
}

type JazzCashInquiryResponse struct {
	ResponseCode           string `json:"pp_ResponseCode"`
	ResponseMessage        string `json:"pp_ResponseMessage"`
	PaymentResponseCode    string `json:"pp_PaymentResponseCode"`
	PaymentResponseMessage string `json:"pp_PaymentResponseMessage"`
	Status                 string `json:"pp_Status"`
}

func (r *JazzCashInquiryResponse) GetResponseCode() string {
	return r.ResponseCode
}

func (r *JazzCashInquiryResponse) GetResponseMessage() string {
	return r.ResponseMessage
}

func (r *JazzCashInquiryResponse) GetPaymentResponseCode() string {
	return r.PaymentResponseCode
}

func (r *JazzCashInquiryResponse) GetPaymentResponseMessage() string {
	return r.PaymentResponseMessage
}

func (r *JazzCashInquiryResponse) GetStatus() string {
	return r.Status
}

type JazzCashResponsePayload struct {
	TxnType              string `json:"pp_TxnType"`
	Version              string `json:"pp_Version"`
	Amount               string `json:"pp_Amount"`
	AuthCode             string `json:"pp_AuthCode"`
	BillReference        string `json:"pp_BillReference"`
	Language             string `json:"pp_Language"`
	MerchantID           string `json:"pp_MerchantID"`
	ResponseCode         string `json:"pp_ResponseCode"`
	ResponseMessage      string `json:"pp_ResponseMessage"`
	RetreivalReferenceNo string `json:"pp_RetreivalReferenceNo"`
	SubMerchantID        string `json:"pp_SubMerchantID"`
	TxnCurrency          string `json:"pp_TxnCurrency"`
	TxnDateTime          string `json:"pp_TxnDateTime"`
	TxnRefNo             string `json:"pp_TxnRefNo"`
	MobileNumber         string `json:"pp_MobileNumber"`
	CNIC                 string `json:"pp_CNIC"`
	DiscountedAmount     string `json:"pp_DiscountedAmount"`
	SecureHash           string `json:"pp_SecureHash"`
}

func (r *JazzCashResponsePayload) GetResponseCode() string {
	return r.ResponseCode
}

func (r *JazzCashResponsePayload) GetResponseMessage() string {
	return r.ResponseMessage
}

func (r *JazzCashResponsePayload) IsValid() bool {
	return r.GetResponseCode() == "000"
}

type NotifJazzCashRequestBody struct {
}

type MidtransPaymentRequestBody struct {
	ReqTransaction struct {
		OrderId     string `json:"order_id"`
		GrossAmount int    `json:"gross_amount"`
	} `json:"transaction_details"`
	ReqCustomer struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
	} `json:"customer_details"`
	ReqCallback struct {
		Finish string `json:"finish"`
	} `json:"callbacks"`
}

type MidtransResponsePayload struct {
	Token       string `json:"token"`
	RedirectUrl string `json:"redirect_url"`
}

func (r *MidtransResponsePayload) GetToken() string {
	return r.Token
}

func (r *MidtransResponsePayload) GetRedirectUrl() string {
	return r.RedirectUrl
}

func (r *MidtransResponsePayload) IsValid() bool {
	return r.RedirectUrl != ""
}

func (r *MidtransResponsePayload) SetRedirectUrl(param string) {
	r.RedirectUrl = r.RedirectUrl + "#/" + param
}

type NotifMidtransRequestBody struct {
	TransactionStatus string `json:"transaction_status"`
	TransactionId     string `json:"transaction_id"`
	OrderId           string `json:"order_id"`
	FraudStatus       string `json:"fraud_status"`
}

func (r *NotifMidtransRequestBody) GetTransactionStatus() string {
	return r.TransactionStatus
}

func (r *NotifMidtransRequestBody) GetTransactionId() string {
	return r.TransactionId
}

func (r *NotifMidtransRequestBody) GetOrderId() string {
	return r.OrderId
}

func (r *NotifMidtransRequestBody) GetFraudStatus() string {
	return r.FraudStatus
}

func (r *NotifMidtransRequestBody) IsValid() bool {
	return r.GetTransactionStatus() == "settlement" || r.GetTransactionStatus() == "capture"
}

type MomoRequestBody struct {
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

func (r *MomoRequestBody) GetPartnerCode() string {
	return r.PartnerCode
}

func (r *MomoRequestBody) GetPartnerName() string {
	return r.PartnerName
}

func (r *MomoRequestBody) GetStoreId() string {
	return r.StoreId
}

func (r *MomoRequestBody) GetRequestId() string {
	return r.RequestId
}

func (r *MomoRequestBody) GetAmount() int {
	return r.Amount
}

func (r *MomoRequestBody) GetOrderId() string {
	return r.OrderId
}

func (r *MomoRequestBody) GetOrderInfo() string {
	return r.OrderInfo
}

func (r *MomoRequestBody) GetRedirectUrl() string {
	return r.RedirectUrl
}

func (r *MomoRequestBody) GetIpnUrl() string {
	return r.IpnUrl
}

func (r *MomoRequestBody) GetRequestType() string {
	return r.RequestType
}

func (r *MomoRequestBody) GetExtraData() string {
	return r.ExtraData
}

func (r *MomoRequestBody) GetLang() string {
	return r.Lang
}

func (r *MomoRequestBody) GetAutoCapture() bool {
	return r.AutoCapture
}

func (r *MomoRequestBody) GetSignature() string {
	return r.Signature
}

type MomoRefundRequestBody struct {
	PartnerCode string `json:"partnerCode"`
	OrderId     string `json:"orderId"`
	RequestId   string `json:"requestId"`
	Amount      int    `json:"amount"`
	TransId     int    `json:"transId"`
	Lang        string `json:"lang"`
	Description string `json:"description"`
	Signature   string `json:"signature"`
}

func (r *MomoRefundRequestBody) GetPartnerCode() string {
	return r.PartnerCode
}

func (r *MomoRefundRequestBody) GetOrderId() string {
	return r.OrderId
}

func (r *MomoRefundRequestBody) GetRequestId() string {
	return r.RequestId
}

func (r *MomoRefundRequestBody) GetAmount() int {
	return r.Amount
}

func (r *MomoRefundRequestBody) GetTransId() int {
	return r.TransId
}

func (r *MomoRefundRequestBody) GetLang() string {
	return r.Lang
}

func (r *MomoRefundRequestBody) GetDescription() string {
	return r.Description
}

func (r *MomoRefundRequestBody) GetSignature() string {
	return r.Signature
}

type MomoResponsePayload struct {
	PartnerCode  string `json:"partnerCode"`
	OrderId      string `json:"orderId"`
	Amount       int    `json:"amount"`
	ResponseTime int    `json:"responseTime"`
	Message      string `json:"message"`
	ResultCode   int    `json:"resultCode"`
	PayUrl       string `json:"payUrl"`
}

func (r *MomoResponsePayload) GetPartnerCode() string {
	return r.PartnerCode
}

func (r *MomoResponsePayload) GetMessage() string {
	return r.Message
}

func (r *MomoResponsePayload) GetResultCode() int {
	return r.ResultCode
}

func (r *MomoResponsePayload) GetPayUrl() string {
	return r.PayUrl
}

func (r *MomoResponsePayload) IsValid() bool {
	return r.GetResultCode() == 0
}

type NotifMomoRequestBody struct {
	PartnerCode  string `json:"partnerCode" query:"partnerCode"`
	OrderId      string `json:"orderId" query:"orderId"`
	RequestId    string `json:"requestId" query:"requestId"`
	Amount       int    `json:"amount" query:"amount"`
	OrderInfo    string `json:"orderInfo" query:"orderInfo"`
	OrderType    string `json:"orderType" query:"orderType"`
	TransId      int    `json:"transId" query:"transId"`
	ResultCode   int    `json:"resultCode" query:"resultCode"`
	Message      string `json:"message" query:"message"`
	PayType      string `json:"payType" query:"payType"`
	ResponseTime int    `json:"responseTime" query:"responseTime"`
	ExtraData    string `json:"extraData" query:"extraData"`
	Signature    string `json:"signature" query:"signature"`
}

func (r *NotifMomoRequestBody) GetPartnerCode() string {
	return r.PartnerCode
}

func (r *NotifMomoRequestBody) GetOrderId() string {
	return r.OrderId
}

func (r *NotifMomoRequestBody) GetRequestId() string {
	return r.RequestId
}

func (r *NotifMomoRequestBody) GetAmount() int {
	return r.Amount
}

func (r *NotifMomoRequestBody) GetOrderInfo() string {
	return r.OrderInfo
}

func (r *NotifMomoRequestBody) GetOrderType() string {
	return r.OrderType
}

func (r *NotifMomoRequestBody) GetTransId() int {
	return r.TransId
}

func (r *NotifMomoRequestBody) GetResultCode() int {
	return r.ResultCode
}

func (r *NotifMomoRequestBody) GetMessage() string {
	return r.Message
}

func (r *NotifMomoRequestBody) GetPayType() string {
	return r.PayType
}

func (r *NotifMomoRequestBody) GetResponseTime() int {
	return r.ResponseTime
}

func (r *NotifMomoRequestBody) GetExtraData() string {
	return r.ExtraData
}

func (r *NotifMomoRequestBody) GetSignature() string {
	return r.Signature
}

func (r *NotifMomoRequestBody) IsValid() bool {
	return r.GetResultCode() == 0
}

type NicepayRequestBody struct {
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
	CartData          string `json:"cartData"`
}

type NicepayRequestBodyItem struct {
	GoodsName     string `json:"goods_name"`
	GoodsDetail   string `json:"goods_detail"`
	GoodsAmt      int    `json:"goods_amt"`
	GoodsQuantity int    `json:"goods_quantity"`
	ImgUrl        string `json:"img_url"`
}

func (r *NicepayRequestBody) GetTimeStamp() string {
	return r.TimeStamp
}

func (r *NicepayRequestBody) GetMerchantId() string {
	return r.MerchantId
}

func (r *NicepayRequestBody) GetPaymentMethod() string {
	return r.PaymentMethod
}

func (r *NicepayRequestBody) GetMitraCode() string {
	return r.MitraCode
}

func (r *NicepayRequestBody) GetCurrency() string {
	return r.Currency
}

func (r *NicepayRequestBody) GetPaymentAmount() string {
	return r.PaymentAmount
}

func (r *NicepayRequestBody) GetReferenceNo() string {
	return r.ReferenceNo
}

func (r *NicepayRequestBody) GetGoodsName() string {
	return r.GoodsName
}

func (r *NicepayRequestBody) GetBuyerName() string {
	return r.BuyerName
}

func (r *NicepayRequestBody) GetBuyerPhone() string {
	return r.BuyerPhone
}

func (r *NicepayRequestBody) GetBuyerEmail() string {
	return r.BuyerEmail
}

func (r *NicepayRequestBody) GetBuyerAddress() string {
	return r.BuyerAddress
}

func (r *NicepayRequestBody) GetBuyerCity() string {
	return r.BuyerCity
}

func (r *NicepayRequestBody) GetBillingState() string {
	return r.BillingState
}

func (r *NicepayRequestBody) GetBillingPostNumber() string {
	return r.BillingPostNumber
}

func (r *NicepayRequestBody) GetBillingCountry() string {
	return r.BillingCountry
}

func (r *NicepayRequestBody) GetNotificationUrl() string {
	return r.NotificationUrl
}

func (r *NicepayRequestBody) GetMerchantToken() string {
	return r.MerchantToken
}

type NicepayResponsePayload struct {
	TransactionId string `json:"tXid"`
}

func (r *NicepayResponsePayload) GetTransactionId() string {
	return r.TransactionId
}

func (r *NicepayResponsePayload) IsValid() bool {
	return r.GetTransactionId() != ""
}

type NotifNicepayRequestBody struct {
	TransactionId   string `json:"tXid" form:"tXid"`
	ReferenceNo     string `json:"referenceNo" form:"referenceNo"`
	PaymentMethod   string `json:"payMethod" form:"payMethod"`
	PaymentAmount   string `json:"amt" form:"amt"`
	TransactionDate string `json:"transDt" form:"transDt"`
	TransactionTime string `json:"transTm" form:"transTm"`
	Currency        string `json:"currency" form:"currency"`
	GoodsName       string `json:"goodsNm" form:"goodsNm"`
	BuyerName       string `json:"billingNm" form:"billingNm"`
	MatchFlag       string `json:"matchCl" form:"matchCl"`
	Status          string `json:"status" form:"status"`
	MerchantToken   string `json:"merchantToken" form:"merchantToken"`
	MitraCode       string `json:"mitraCd" form:"mitraCd"`
}

func (r *NotifNicepayRequestBody) GetTransactionId() string {
	return r.TransactionId
}

func (r *NotifNicepayRequestBody) GetReferenceNo() string {
	return r.ReferenceNo
}

func (r *NotifNicepayRequestBody) GetPaymentMethod() string {
	return r.PaymentMethod
}

func (r *NotifNicepayRequestBody) GetPaymentAmount() string {
	return r.PaymentAmount
}

func (r *NotifNicepayRequestBody) GetTransactionDate() string {
	return r.TransactionDate
}

func (r *NotifNicepayRequestBody) GetTransactionTime() string {
	return r.TransactionTime
}

func (r *NotifNicepayRequestBody) GetCurrency() string {
	return r.Currency
}

func (r *NotifNicepayRequestBody) GetGoodsName() string {
	return r.GoodsName
}

func (r *NotifNicepayRequestBody) GetBuyerName() string {
	return r.BuyerName
}

func (r *NotifNicepayRequestBody) GetMatchFlag() string {
	return r.MatchFlag
}

func (r *NotifNicepayRequestBody) GetStatus() string {
	return r.Status
}

func (r *NotifNicepayRequestBody) GetMerchantToken() string {
	return r.MerchantToken
}

func (r *NotifNicepayRequestBody) GetMitraCode() string {
	return r.MitraCode
}

func (r *NotifNicepayRequestBody) IsValid() bool {
	return r.GetStatus() == "0"
}

type RazerRequestBody struct {
	MerchantId string `form:"merchant_id"`
	Amount     int    `form:"amount"`
	OrderId    string `form:"orderid"`
	BillName   string `form:"bill_name"`
	BillEmail  string `form:"bill_email"`
	BillMobile string `form:"bill_mobile"`
	BillDesc   string `form:"bill_desc"`
	Vcode      string `form:"vcode"`
}

func (r *RazerRequestBody) GetMerchantId() string {
	return r.MerchantId
}

func (r *RazerRequestBody) GetAmount() int {
	return r.Amount
}

func (r *RazerRequestBody) GetOrderId() string {
	return r.OrderId
}

func (r *RazerRequestBody) GetBillName() string {
	return r.BillName
}

func (r *RazerRequestBody) GetBillEmail() string {
	return r.BillEmail
}

func (r *RazerRequestBody) GetBillMobile() string {
	return r.BillMobile
}

func (r *RazerRequestBody) GetBillDesc() string {
	return r.BillDesc
}

func (r *RazerRequestBody) GetVcode() string {
	return r.Vcode
}

type RazerResponsePayload struct {
	RedirectUrl string `json:"redirect_url"`
}

func (r *RazerResponsePayload) GetRedirectUrl() string {
	return r.RedirectUrl
}

type NotifRazerRequestBody struct {
	TranId   string `json:"tranID" form:"tranID"`
	OrderId  string `json:"orderid" form:"orderid"`
	Status   string `json:"status" form:"status"`
	Domain   string `json:"domain" form:"domain"`
	Amount   string `json:"amount" form:"amount"`
	Currency string `json:"currency" form:"currency"`
	AppCode  string `json:"appcode" form:"appcode"`
	PayDate  string `json:"paydate" form:"paydate"`
	Skey     string `json:"skey" form:"skey"`
}

func (r *NotifRazerRequestBody) GetTranId() string {
	return r.TranId
}

func (r *NotifRazerRequestBody) GetOrderId() string {
	return r.OrderId
}

func (r *NotifRazerRequestBody) GetStatus() string {
	return r.Status
}

func (r *NotifRazerRequestBody) GetDomain() string {
	return r.Domain
}

func (r *NotifRazerRequestBody) GetAmount() string {
	return r.Amount
}

func (r *NotifRazerRequestBody) GetCurrency() string {
	return r.Currency
}

func (r *NotifRazerRequestBody) GetAppCode() string {
	return r.AppCode
}

func (r *NotifRazerRequestBody) GetPayDate() string {
	return r.PayDate
}

func (r *NotifRazerRequestBody) GetSkey() string {
	return r.Skey
}

func (r *NotifRazerRequestBody) IsValid() bool {
	return r.GetStatus() == "00"
}

type RefundRazerRequestBody struct {
	TransactionId string `json:"TxnID" form:"TxnID"`
	MerchantID    string `json:"MerchantID" form:"MerchantID"`
	Signature     string `json:"Signature" form:"Signature"`
}

func (r *RefundRazerRequestBody) GetTransactionId() string {
	return r.TransactionId
}

func (r *RefundRazerRequestBody) GetMerchantId() string {
	return r.MerchantID
}

func (r *RefundRazerRequestBody) GetSignature() string {
	return r.Signature
}

func (r *RefundRazerRequestBody) SetSignature(data string) {
	r.Signature = data
}

type XimpayTselRequestBody struct {
	PartnerId string `json:"partnerid"`
	ItemId    string `json:"itemid"`
	CbParam   string `json:"cbparam"`
	Token     string `json:"token"`
	Op        string `json:"op"`
	Msisdn    string `json:"msisdn"`
}

func (e *XimpayTselRequestBody) SetItemId(data string) {
	e.ItemId = data
}

func (e *XimpayTselRequestBody) SetToken(data string) {
	e.Token = data
}

type XimpayHtiRequestBody struct {
	PartnerId string `json:"partnerid"`
	ItemId    string `json:"itemid"`
	CbParam   string `json:"cbparam"`
	Token     string `json:"token"`
	Op        string `json:"op"`
	Msisdn    string `json:"msisdn"`
}

func (e *XimpayHtiRequestBody) SetItemId(data string) {
	e.ItemId = data
}

func (e *XimpayHtiRequestBody) SetToken(data string) {
	e.Token = data
}

type XimpayIsatRequestBody struct {
	PartnerId  string `json:"partnerid"`
	ItemName   string `json:"item_name"`
	ItemDesc   string `json:"item_desc"`
	Amount     int    `json:"amount"`
	ChargeType string `json:"charge_type"`
	CbParam    string `json:"cbparam"`
	Token      string `json:"token"`
	Op         string `json:"op"`
	Msisdn     string `json:"msisdn"`
}

type XimpayXlRequestBody struct {
	PartnerId string `json:"partnerid"`
	ItemName  string `json:"item_name"`
	ItemDesc  string `json:"item_desc"`
	Amount    int    `json:"amount"`
	CbParam   string `json:"cbparam"`
	Token     string `json:"token"`
	Op        string `json:"op"`
	Msisdn    string `json:"msisdn"`
}

type XimpaySfRequestBody struct {
	PartnerId string `json:"partnerid"`
	ItemName  string `json:"item_name"`
	ItemDesc  string `json:"item_desc"`
	AmountExc int    `json:"amount_exc"`
	CbParam   string `json:"cbparam"`
	Token     string `json:"token"`
	Op        string `json:"op"`
	Msisdn    string `json:"msisdn"`
}

type XimpayPinRequestBody struct {
	XimpayId    string `json:"ximpayid"`
	CodePin     string `json:"codepin"`
	XimpayToken string `json:"ximpaytoken"`
}

type XimpayTransactionResponse struct {
	Transaction []XimpayResponse `json:"ximpaytransaction"`
}

type XimpayResponse struct {
	ResponseCode int    `json:"responsecode"`
	XimpayId     string `json:"ximpayid"`
}

func (e *XimpayTransactionResponse) IsValid() bool {
	return e.Transaction[0].ResponseCode == 1
}

func (e *XimpayTransactionResponse) IsWrongPhoneNumber() bool {
	return e.Transaction[0].ResponseCode == -9 || e.Transaction[0].ResponseCode == -10
}

func (e *XimpayTransactionResponse) IsWrongPIN() bool {
	return e.Transaction[0].ResponseCode == -2
}

func (e *XimpayTransactionResponse) IsInvalidPIN() bool {
	return e.Transaction[0].ResponseCode == -13
}

func (e *XimpayTransactionResponse) GetXimpayId() string {
	return e.Transaction[0].XimpayId
}

type NotifXimpayRequestParam struct {
	XimpayId     string `query:"ximpayid" json:"ximpayid"`
	XimpayStatus string `query:"ximpaystatus" json:"ximpaystatus"`
	CbParam      string `query:"cbparam" json:"cbparam"`
	XimpayToken  string `query:"ximpaytoken" json:"ximpaytoken"`
	FailCode     string `query:"failcode" json:"failcode"`
}

func (e *NotifXimpayRequestParam) IsValidXimpayToken(secretKey string) bool {
	str := e.GetXimpayId() + e.GetXimpayStatus() + e.GetCbParam() + secretKey
	return e.GetXimpayToken() == hash_utils.GetMD5Hash(strings.ToLower(str))
}

func (e *NotifXimpayRequestParam) GetXimpayId() string {
	return e.XimpayId
}

func (e *NotifXimpayRequestParam) GetCbParam() string {
	return e.CbParam
}

func (e *NotifXimpayRequestParam) GetFailCode() string {
	return e.FailCode
}

func (e *NotifXimpayRequestParam) GetXimpayStatus() string {
	return e.XimpayStatus
}

func (e *NotifXimpayRequestParam) GetXimpayToken() string {
	return e.XimpayToken
}

func (e *NotifXimpayRequestParam) IsValid() bool {
	return e.GetXimpayStatus() == "1" && e.GetFailCode() == "0"
}

type XenditPayoutRequest struct {
	ExternalId         string  `json:"external_id"`
	Amount             float64 `json:"amount"`
	Email              string  `json:"email"`
	InvoiceDuration    int     `json:"invoice_duration"`
	SuccessRedirectUrl string  `json:"success_redirect_url"`
	FailureRedirectUrl string  `json:"failure_redirect_url"`
}

func (e *XenditPayoutRequest) SetExternalId(v string) {
	e.ExternalId = v
}

func (e *XenditPayoutRequest) SetAmount(v float64) {
	e.Amount = v
}

func (e *XenditPayoutRequest) SetInvoiceDuration(v int) {
	e.InvoiceDuration = v
}

func (e *XenditPayoutRequest) SetSuccessRedirectUrl(v string) {
	e.SuccessRedirectUrl = v
}

func (e *XenditPayoutRequest) SetFailureRedirectUrl(v string) {
	e.FailureRedirectUrl = v
}

type XenditPayoutResponse struct {
	Id           string `json:"id"`
	ExternalId   string `json:"external_id"`
	Amount       int    `json:"amount"`
	MerchantName string `json:"merchant_name"`
	Status       string `json:"status"`
	ExpiryDate   string `json:"expiry_date"`
	Created      string `json:"created"`
	InvoiceUrl   string `json:"invoice_url"`
	Currency     string `json:"currency"`
}

func (e *XenditPayoutResponse) GetExternalId() string {
	return e.ExternalId
}

func (e *XenditPayoutResponse) GetStatus() string {
	return e.Status
}

func (e *XenditPayoutResponse) GetInvoiceUrl() string {
	return e.InvoiceUrl
}

type NotifXenditRequestBody struct {
	ExternalId string `json:"external_id"`
	Status     string `json:"status"`
	Amount     int    `json:"amount"`
	PaidAmount int    `json:"paid_amount"`
	Currency   string `json:"currency"`
}

func (e *NotifXenditRequestBody) GetExternalId() string {
	return e.ExternalId
}

func (e *NotifXenditRequestBody) IsValid() bool {
	return e.Status == "PAID"
}

type CallbackRequestBody struct {
	Number string    `json:"number"`
	IsPaid bool      `json:"is_paid"`
	Time   time.Time `json:"time"`
}

func (r *CallbackRequestBody) GetNumber() string {
	return r.Number
}

func (r *CallbackRequestBody) GetIsPaid() bool {
	return r.IsPaid
}

func (r *CallbackRequestBody) GetTime() time.Time {
	return r.Time
}

type NotifRequestBody struct {
	NotifMidtransRequestBody  *NotifMidtransRequestBody
	NotifNicepayRequestBody   *NotifNicepayRequestBody
	NotifDragonPayRequestBody *NotifDragonPayRequestBody
	NotifJazzCashRequestBody  *NotifJazzCashRequestBody
	NotifMomoRequestBody      *NotifMomoRequestBody
	NotifRazerRequestBody     *NotifRazerRequestBody
	NotifXimpayRequestBody    *NotifXimpayRequestParam
	NotifXenditRequestBody    *NotifXenditRequestBody
	Channel                   string `json:"channel"`
}

func (r *NotifRequestBody) GetChannel() string {
	return strings.ToUpper(r.Channel)
}

func (r *NotifRequestBody) SetChannel(data string) {
	r.Channel = strings.ToUpper(data)
}

func (r *NotifRequestBody) IsMidtrans() bool {
	return r.GetChannel() == "MIDTRANS"
}

func (r *NotifRequestBody) IsNicePay() bool {
	return r.GetChannel() == "NICEPAY"
}

func (r *NotifRequestBody) IsDragonPay() bool {
	return r.GetChannel() == "DRAGONPAY"
}

func (r *NotifRequestBody) IsJazzCash() bool {
	return r.GetChannel() == "JAZZCASH"
}

func (r *NotifRequestBody) IsMomo() bool {
	return r.GetChannel() == "MOMO"
}

func (r *NotifRequestBody) IsRazer() bool {
	return r.GetChannel() == "RAZER"
}

func (r *NotifRequestBody) IsXimpay() bool {
	return r.GetChannel() == "XIMPAY"
}

func (r *NotifRequestBody) IsXendit() bool {
	return r.GetChannel() == "XENDIT"
}
