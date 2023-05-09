package entity

type OrderRequestBody struct {
	Msisdn string `json:"msisdn"`
	Email  string `json:"email"`
	Number string `json:"number"`
}

type DragonPayRequestBody struct {
}

type DragonPayResponsePayload struct {
}

type NotifDragonPayRequestBody struct {
}

type JazzCashRequestBody struct {
}

type JazzCashResponsePayload struct {
}

type NotifJazzCashRequestBody struct {
}

type MidtransRequestBody struct {
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
