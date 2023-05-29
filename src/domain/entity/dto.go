package entity

import "time"

type OrderRequestBody struct {
	Channel     string  `json:"channel"`
	Msisdn      string  `json:"msisdn"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Number      string  `json:"number"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	UrlCallback string  `json:"url_callback"`
	IpAddress   string  `json:"ip_address"`
}

func (r *OrderRequestBody) GetChannel() string {
	return r.Channel
}

func (r *OrderRequestBody) GetMsisdn() string {
	return r.Msisdn
}

func (r *OrderRequestBody) GetName() string {
	return r.Name
}

func (r *OrderRequestBody) GetEmail() string {
	return r.Email
}

func (r *OrderRequestBody) GetNumber() string {
	return r.Number
}

func (r *OrderRequestBody) GetAmount() float64 {
	return r.Amount
}

func (r *OrderRequestBody) GetDescription() string {
	return r.Description
}

func (r *OrderRequestBody) GetUrlCallback() string {
	return r.UrlCallback
}

func (r *OrderRequestBody) GetIpAddress() string {
	return r.IpAddress
}

type RefundRequestBody struct {
	Number      string  `json:"number"`
	Amount      float64 `json:"amount"`
	UrlCallback string  `json:"url_callback"`
}

func (r *RefundRequestBody) GetNumber() string {
	return r.Number
}

func (r *RefundRequestBody) GetAmount() float64 {
	return r.Amount
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
	// not yet
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
	CartData          struct {
		Count                  string   `json:"count"`
		NicepayRequestBodyItem []string `json:"item"`
	} `json:"cartData"`
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
