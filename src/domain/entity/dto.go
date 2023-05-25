package entity

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
	AppCode string  `json:"app_code"`
	Number  string  `json:"number"`
	Amount  float64 `json:"amount"`
}

type DragonPayRequestBody struct {
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

type JazzCashRefundRequest struct {
	TxnRefNo     string `json:"pp_TxnRefNo"`
	Amount       string `json:"pp_Amount"`
	TxnCurrency  string `json:"pp_TxnCurrency"`
	MerchantID   string `json:"pp_MerchantID"`
	Password     string `json:"pp_Password"`
	MerchantMPIN string `json:"pp_MerchantMPIN"`
	SecureHash   string `json:"pp_SecureHash"`
}

type JazzCashRefundResponse struct {
	SecureHash      string `json:"pp_SecureHash"`
	ResponseCode    string `json:"pp_ResponseCode"`
	ResponseMessage string `json:"pp_ResponseMessage"`
}

type JazzCashInquiryRequest struct {
	TxnRefNo   string `json:"pp_TxnRefNo"`
	MerchantID string `json:"pp_MerchantID"`
	Password   string `json:"pp_Password"`
	SecureHash string `json:"pp_SecureHash"`
}

type JazzCashInquiryResponse struct {
	ResponseCode           string `json:"pp_ResponseCode"`
	ResponseMessage        string `json:"pp_ResponseMessage"`
	PaymentResponseCode    string `json:"pp_PaymentResponseCode"`
	PaymentResponseMessage string `json:"pp_PaymentResponseMessage"`
	Status                 string `json:"pp_Status"`
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

type MomoResponsePayload struct {
	PartnerCode  string `json:"partnerCode"`
	OrderId      string `json:"orderId"`
	Amount       int    `json:"amount"`
	ResponseTime int    `json:"responseTime"`
	Message      string `json:"message"`
	ResultCode   int    `json:"resultCode"`
	PayUrl       string `json:"payUrl"`
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
}

type NicepayRequestBody struct {
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

type RazerResponsePayload struct {
}

type NotifRazerRequestBody struct {
}

type PostbackRequestBody struct {
	Msisdn string `json:"msisdn"`
	Number string `json:"number"`
}
