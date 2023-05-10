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

	// init order
	orderRepo := repository.NewOrderRepository(u.db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(u.cfg, orderService)

	// init payment
	paymentRepo := repository.NewPaymentRepository(u.db)
	paymentService := services.NewPaymentService(paymentRepo)
	notifyHandler := handler.NewNotifyHandler(u.cfg, orderService, paymentService)

	// init base
	baseHandler := handler.NewBaseHandler(u.cfg)

	router.Static("/static", path+"/public")

	/**
	 * Routes Base
	 */
	router.Get("/", baseHandler.Base)

	/**
	 * Routes Order & Notify
	 */
	dragopay := router.Group("dragonpay")
	dragopay.Post("order", orderHandler.DragonPay)
	dragopay.Post("notification", notifyHandler.DragonPay)

	jazzcash := router.Group("jazzcash")
	jazzcash.Post("order", orderHandler.JazzCash)
	jazzcash.Post("notification", notifyHandler.JazzCash)

	midtrans := router.Group("midtrans")
	midtrans.Post("order", orderHandler.Midtrans)
	midtrans.Post("notification", notifyHandler.Midtrans)

	momo := router.Group("momo")
	momo.Post("order", orderHandler.Momo)
	momo.Post("notification", notifyHandler.Momo)

	nicepay := router.Group("nicepay")
	nicepay.Post("order", orderHandler.Nicepay)
	nicepay.Post("notification", notifyHandler.Nicepay)

	razer := router.Group("razer")
	razer.Post("order", orderHandler.Razer)
	razer.Post("notification", notifyHandler.Razer)

	return router
}
