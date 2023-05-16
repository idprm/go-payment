package app

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/idprm/go-payment/src/domain/repository"
	"github.com/idprm/go-payment/src/handler"
	"github.com/idprm/go-payment/src/services"
)

func (u *UrlMappings) mapUrls() *fiber.App {
	/**
	 * Init Fiber
	 */
	router := fiber.New()

	/**
	 * Access log on browser
	 */
	router.Use("/logs", filesystem.New(filesystem.Config{
		Root:         http.Dir(u.cfg.Log.Path),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
		MaxAge:       3600,
	}))

	/**
	 * Write access logger
	 */
	// file, err := os.OpenFile(u.cfg.Log.Path+"/access_log/log-"+time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }

	// router.Use(requestid.New())
	// router.Use(log_access.New(log_access.Config{
	// 	Format:     "${pid} ${status} - ${method} ${path}\n",
	// 	TimeFormat: "02-Jan-2006",
	// 	TimeZone:   u.cfg.App.TimeZone,
	// 	Output:     file,
	// }))

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	router.Static("/static", path+"/public")

	// init transaction
	transactionRepo := repository.NewTransactionRepository(u.db)
	transactionService := services.NewTransactionService(transactionRepo)

	// init order
	orderRepo := repository.NewOrderRepository(u.db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(u.cfg, orderService, transactionService)

	// init callback
	callbackRepo := repository.NewCallbackRepository(u.db)
	callbackService := services.NewCallbackService(callbackRepo)

	// init payment
	paymentRepo := repository.NewPaymentRepository(u.db)
	paymentService := services.NewPaymentService(orderRepo, paymentRepo)
	paymentHandler := handler.NewPaymentHandler(u.cfg, paymentService, transactionService, callbackService)

	// init refund
	refundRepo := repository.NewRefundRepository(u.db)
	refundService := services.NewRefundService(orderRepo, refundRepo)
	refundHandler := handler.NewRefundHandler(u.cfg, refundService, transactionService)

	// init base
	baseHandler := handler.NewBaseHandler(u.cfg)

	/**
	 * Routes Base
	 */
	router.Get("/", baseHandler.Base)

	authenticated := router.Group("auth")

	/**
	 * Routes Order & Notify (version 1)
	 */
	v1 := router.Group("v1")

	dragopay := v1.Group("dragonpay")
	dragopay.Post("order", orderHandler.DragonPay)
	dragopay.Post("notification", paymentHandler.DragonPay)
	dragopay.Post("refund", refundHandler.DragonPay)

	jazzcash := v1.Group("jazzcash")
	jazzcash.Post("order", orderHandler.JazzCash)
	jazzcash.Post("notification", paymentHandler.JazzCash)
	jazzcash.Post("refund", refundHandler.JazzCash)

	midtrans := v1.Group("midtrans")
	midtrans.Post("order", orderHandler.Midtrans)
	midtrans.Post("notification", paymentHandler.Midtrans)
	midtrans.Post("refund", refundHandler.Midtrans)

	momo := v1.Group("momo")
	momo.Post("order", orderHandler.Momo)
	momo.Post("notification", paymentHandler.Momo)
	momo.Post("refund", refundHandler.Momo)

	nicepay := v1.Group("nicepay")
	nicepay.Post("order", orderHandler.Nicepay)
	nicepay.Post("notification", paymentHandler.Nicepay)
	nicepay.Post("refund", refundHandler.Nicepay)

	razer := v1.Group("razer")
	razer.Post("order", orderHandler.Razer)
	razer.Post("notification", paymentHandler.Razer)
	razer.Post("refund", refundHandler.Razer)

	/**
	 * AUTHENTICATED ROUTING
	 */
	authOrders := authenticated.Group("orders")
	authOrders.Get("/", orderHandler.GetAll)
	authOrders.Get("/:id", orderHandler.Get)
	authOrders.Put("/", orderHandler.Update)
	authOrders.Delete("/:id", orderHandler.Delete)

	authPayments := authenticated.Group("payments")
	authPayments.Get("/", paymentHandler.GetAll)
	authPayments.Get("/:id", paymentHandler.Get)
	authPayments.Put("/", paymentHandler.Update)
	authPayments.Delete("/:id", paymentHandler.Delete)

	authRefunds := authenticated.Group("refunds")
	authRefunds.Get("/", refundHandler.GetAll)
	authRefunds.Get("/:id", refundHandler.Get)
	authRefunds.Put("/", refundHandler.Update)
	authRefunds.Delete("/:id", refundHandler.Delete)

	return router
}
