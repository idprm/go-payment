package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	log_access "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/idprm/go-payment/src/domain/repository"
	"github.com/idprm/go-payment/src/handler"
	"github.com/idprm/go-payment/src/services"
)

func (u *UrlMappings) mapUrls() *fiber.App {
	/**
	 * Init Fiber
	 */
	router := fiber.New()

	// set cors
	router.Use(cors.New())

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
	file, err := os.OpenFile(u.cfg.Log.Path+"/access_log/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	router.Use(requestid.New())
	router.Use(log_access.New(log_access.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   u.cfg.App.TimeZone,
		Output:     file,
	}))

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	router.Static("/static", path+"/public")

	// init country
	countryRepo := repository.NewCountryRepository(u.db)
	countryService := services.NewCountryService(countryRepo)

	// init gateway
	gatewayRepo := repository.NewGatewayRepository(u.db)
	gatewayService := services.NewGatewayService(gatewayRepo)

	// init channel
	channelRepo := repository.NewChannelRepository(u.db)
	channelService := services.NewChannelService(u.cfg, gatewayRepo, channelRepo)

	gatewayHandler := handler.NewGatewayHandler(u.cfg, countryService, gatewayService, channelService)
	channelHandler := handler.NewChannelHandler(u.cfg, gatewayService, channelService)

	// init application
	applicationRepo := repository.NewApplicationRepository(u.db)
	applicationService := services.NewApplicationService(applicationRepo)

	// init transaction
	transactionRepo := repository.NewTransactionRepository(u.db)
	transactionService := services.NewTransactionService(transactionRepo)

	// init order
	orderRepo := repository.NewOrderRepository(u.db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(u.cfg, u.logger, u.zap, applicationService, channelService, orderService, transactionService)

	// init callback
	callbackRepo := repository.NewCallbackRepository(u.db)
	callbackService := services.NewCallbackService(callbackRepo)

	// init payment
	paymentRepo := repository.NewPaymentRepository(u.db)
	paymentService := services.NewPaymentService(orderRepo, paymentRepo)
	paymentHandler := handler.NewPaymentHandler(u.cfg, u.rds, u.logger, u.zap, orderService, paymentService, transactionService, callbackService, u.ctx)

	// init refund
	refundRepo := repository.NewRefundRepository(u.db)
	refundService := services.NewRefundService(orderRepo, refundRepo)
	refundHandler := handler.NewRefundHandler(u.cfg, u.logger, u.zap, applicationService, orderService, paymentService, refundService, transactionService)

	// init return
	returnRepo := repository.NewReturnRepository(u.db)
	returnService := services.NewReturnService(orderRepo, returnRepo)
	returnHandler := handler.NewReturnHandler(u.cfg, u.logger, u.zap, orderService, transactionService, returnService)

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

	country := router.Group("country")
	country.Get("/", gatewayHandler.Country)
	country.Get("/:locale", gatewayHandler.Locale)

	dragopay := v1.Group("dragonpay")
	dragopay.Get("/", gatewayHandler.Dragonpay)
	dragopay.Get("channel/:slug", channelHandler.Dragonpay)
	dragopay.Post("order", orderHandler.DragonPay)
	dragopay.Post("notification", paymentHandler.DragonPay)
	dragopay.Post("refund", refundHandler.DragonPay)

	jazzcash := v1.Group("jazzcash")
	jazzcash.Get("/", gatewayHandler.Jazzcash)
	jazzcash.Get("channel/:slug", channelHandler.Jazzcash)
	jazzcash.Post("order", orderHandler.JazzCash)
	jazzcash.Post("notification", paymentHandler.JazzCash)
	jazzcash.Post("refund", refundHandler.JazzCash)

	midtrans := v1.Group("midtrans")
	midtrans.Get("/", gatewayHandler.Midtrans)
	midtrans.Get("channel/:slug", channelHandler.Midtrans)
	midtrans.Post("order", orderHandler.Midtrans)
	midtrans.Post("notification", paymentHandler.Midtrans)
	midtrans.Post("refund", refundHandler.Midtrans)

	momo := v1.Group("momo")
	momo.Get("/", gatewayHandler.Momo)
	momo.Get("channel/:slug", channelHandler.Momo)
	momo.Post("order", orderHandler.Momo)
	momo.Post("notification", paymentHandler.Momo)
	momo.Post("refund", refundHandler.Momo)

	nicepay := v1.Group("nicepay")
	nicepay.Get("/", gatewayHandler.Nicepay)
	nicepay.Get("channel/:slug", channelHandler.Nicepay)
	nicepay.Post("order", orderHandler.Nicepay)
	nicepay.Post("notification", paymentHandler.Nicepay)
	nicepay.Post("refund", refundHandler.Nicepay)

	razer := v1.Group("razer")
	razer.Get("/", gatewayHandler.Razer)
	razer.Get("channel/:slug", channelHandler.Razer)
	razer.Post("order", orderHandler.Razer)
	razer.Post("notification", paymentHandler.Razer)
	razer.Post("refund", refundHandler.Razer)
	razer.Post("return", returnHandler.Razer)

	ximpay := v1.Group("ximpay")
	ximpay.Get("/", gatewayHandler.Ximpay)
	ximpay.Get("channel/:slug", channelHandler.Ximpay)
	ximpay.Post("order", orderHandler.Ximpay)
	ximpay.Get("notification", paymentHandler.Ximpay)

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
