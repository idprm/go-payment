package entity

type OrderRequestBody struct {
	Msisdn string  `json:"msisdn"`
	Email  string  `json:"email"`
	Number string  `json:"number"`
	Amount float64 `json:"amount"`
}

type RefundRequestBody struct {
	Number string  `json:"number"`
	Amount float64 `json:"amount"`
}

type DragonPayRequestBody struct {
}

type DragonPayResponsePayload struct {
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
}

type NotifMidtransRequestBody struct {
}

type MomoRequestBody struct {
}

type MomoResponsePayload struct {
}

type NotifMomoRequestBody struct {
}

type NicepayRequestBody struct {
}

type NicepayResponsePayload struct {
}

type NotifNicepayRequestBody struct {
}

type RazerRequestBody struct {
}

type RazerResponsePayload struct {
}

type NotifRazerRequestBody struct {
}

type PostbackRequestBody struct {
	Msisdn string `json:"msisdn"`
	Number string `json:"number"`
}
