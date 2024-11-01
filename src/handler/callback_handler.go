package handler

import (
	"encoding/json"

	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/providers/callback"
	"github.com/idprm/go-payment/src/services"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type CallbackHandler struct {
	rds                *redis.Client
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	orderService       services.IOrderService
	paymentService     services.IPaymentService
	transactionService services.ITransactionService
	callbackService    services.ICallbackService
}

func NewCallbackHandler(
	rds *redis.Client,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	orderService services.IOrderService,
	paymentService services.IPaymentService,
	transactionService services.ITransactionService,
	callbackService services.ICallbackService,
) *CallbackHandler {
	return &CallbackHandler{
		rds:                rds,
		logger:             logger,
		zap:                zap,
		orderService:       orderService,
		paymentService:     paymentService,
		transactionService: transactionService,
		callbackService:    callbackService,
	}
}

func (h *CallbackHandler) Midtrans(req *entity.NotifMidtransRequestBody) {
	// get order
	order, err := h.orderService.GetByNumber(req.GetOrderId())
	if err != nil {
		h.zap.Error(err)
	}

	if req.IsValid() {
		if !h.paymentService.CountByOrderId(int(order.GetId())) {
			// insert payment
			payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
			if err != nil {
				h.zap.Error(err)
			}
			// hit callback
			provider := callback.NewCallback(h.logger, order.Application, order)
			cb, err := provider.Hit()
			h.zap.Info(string(cb))
			if err != nil {
				h.zap.Error(err)
			}

			dataJson, _ := json.Marshal(req)
			// insert transaction
			h.transactionService.Save(&entity.Transaction{
				ApplicationID: order.Application.GetId(),
				Action:        PAYMENT + MIDTRANS,
				Payload:       string(dataJson),
			})
			// insert callback
			h.callbackService.Save(&entity.Callback{
				PaymentID: payment.GetId(),
				Payload:   string(cb),
			})

		}

	}
}

func (h *CallbackHandler) Nicepay(req *entity.NotifNicepayRequestBody) {
	// get order
	order, err := h.orderService.GetByNumber(req.GetReferenceNo())
	if err != nil {
		h.zap.Error(err)

	}

	if req.IsValid() {

		if !h.paymentService.CountByOrderId(int(order.GetId())) {
			// insert payment
			payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
			if err != nil {
				h.zap.Error(err)

			}
			// hit callback
			provider := callback.NewCallback(h.logger, order.Application, order)
			cb, err := provider.Hit()
			if err != nil {
				h.zap.Error(err)

			}

			dataJson, _ := json.Marshal(req)
			// insert transaction
			h.transactionService.Save(&entity.Transaction{
				ApplicationID: order.Application.GetId(),
				Action:        PAYMENT + NICEPAY,
				Payload:       string(dataJson),
			})
			// insert callback
			h.callbackService.Save(&entity.Callback{
				PaymentID: payment.GetId(),
				Payload:   string(cb),
			})
		}
	}
}

func (h *CallbackHandler) DragonPay(req *entity.NotifDragonPayRequestBody) {
	// get order
	order, err := h.orderService.GetByNumber(req.GetTransactionId())
	if err != nil {
		h.zap.Error(err)
	}

	if req.IsValid() {
		// hit callback
		if !h.paymentService.CountByOrderId(int(order.GetId())) {
			// insert payment
			payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
			if err != nil {
				h.zap.Error(err)
			}
			provider := callback.NewCallback(h.logger, order.Application, order)
			cb, err := provider.Hit()
			h.zap.Info(string(cb))
			if err != nil {
				h.zap.Error(err)
			}

			dataJson, _ := json.Marshal(req)
			// insert transaction
			h.transactionService.Save(&entity.Transaction{
				ApplicationID: order.Application.GetId(),
				Action:        PAYMENT + DRAGONPAY,
				Payload:       string(dataJson),
			})
			// insert callback
			h.callbackService.Save(&entity.Callback{
				PaymentID: payment.GetId(),
				Payload:   string(cb),
			})

		}
	}
}

func (h *CallbackHandler) JazzCash(req *entity.NotifJazzCashRequestBody) {
}

func (h *CallbackHandler) Momo(req *entity.NotifMomoRequestBody) {
	// get order
	order, err := h.orderService.GetByNumber(req.GetOrderId())
	if err != nil {
		h.zap.Error(err)
	}

	if req.IsValid() {
		if !h.paymentService.CountByOrderId(int(order.GetId())) {
			// insert payment
			payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
			if err != nil {
				h.zap.Error(err)

			}
			// hit callback
			provider := callback.NewCallback(h.logger, order.Application, order)
			cb, err := provider.Hit()
			if err != nil {
				h.zap.Error(err)
			}

			dataJson, _ := json.Marshal(req)
			// insert transaction
			h.transactionService.Save(&entity.Transaction{
				ApplicationID: order.Application.GetId(),
				Action:        PAYMENT + MOMO,
				Payload:       string(dataJson),
			})
			// insert callback
			h.callbackService.Save(&entity.Callback{
				PaymentID: payment.GetId(),
				Payload:   string(cb),
			})
		}
	}
}

func (h *CallbackHandler) Razer(req *entity.NotifRazerRequestBody) {
	// get order
	order, err := h.orderService.GetByNumber(req.GetOrderId())
	if err != nil {
		h.zap.Error(err)
	}

	if req.IsValid() {
		if !h.paymentService.CountByOrderId(int(order.GetId())) {
			// insert payment
			payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
			if err != nil {
				h.zap.Error(err)
			}

			// hit callback
			provider := callback.NewCallback(h.logger, order.Application, order)
			cb, err := provider.Hit()
			if err != nil {
				h.zap.Error(err)
			}

			dataJson, _ := json.Marshal(req)
			// insert transaction
			h.transactionService.Save(&entity.Transaction{
				ApplicationID: order.Application.GetId(),
				Action:        PAYMENT + RAZER,
				Payload:       string(dataJson),
			})
			// insert callback
			h.callbackService.Save(&entity.Callback{
				PaymentID: payment.GetId(),
				Payload:   string(cb),
			})
		}

	}
}

func (h *CallbackHandler) Ximpay(req *entity.NotifXimpayRequestParam) {
	// get order
	order, err := h.orderService.GetByNumber(req.GetCbParam())
	if err != nil {
		h.zap.Error(err)
	}

	if req.IsValid() {

		if !h.paymentService.CountByOrderId(int(order.GetId())) {
			// insert payment
			payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
			if err != nil {
				h.zap.Error(err)
			}

			// hit callback
			provider := callback.NewCallback(h.logger, order.Application, order)
			cb, err := provider.Hit()
			if err != nil {
				h.zap.Error(err)
			}

			dataJson, _ := json.Marshal(req)
			// insert transaction
			h.transactionService.Save(&entity.Transaction{
				ApplicationID: order.Application.GetId(),
				Action:        PAYMENT + XIMPAY,
				Payload:       string(dataJson),
			})
			// insert callback
			h.callbackService.Save(&entity.Callback{
				PaymentID: payment.GetId(),
				Payload:   string(cb),
			})

		}
	}
}

func (h *CallbackHandler) Xendit(req *entity.NotifXenditRequestBody) {
	// get order
	order, err := h.orderService.GetByNumber(req.GetExternalId())
	if err != nil {
		h.zap.Error(err)
	}

	if req.IsValid() {

		if !h.paymentService.CountByOrderId(int(order.GetId())) {
			// insert payment
			payment, err := h.paymentService.Save(&entity.Payment{OrderID: order.GetId(), IsPaid: true})
			if err != nil {
				h.zap.Error(err)
			}

			// hit callback
			provider := callback.NewCallback(h.logger, order.Application, order)
			cb, err := provider.Hit()
			if err != nil {
				h.zap.Error(err)
			}

			dataJson, _ := json.Marshal(req)
			// insert transaction
			h.transactionService.Save(&entity.Transaction{
				ApplicationID: order.Application.GetId(),
				Action:        PAYMENT + XENDIT,
				Payload:       string(dataJson),
			})
			// insert callback
			h.callbackService.Save(&entity.Callback{
				PaymentID: payment.GetId(),
				Payload:   string(cb),
			})

		}
	}
}
