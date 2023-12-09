package cmd

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/datasource/mysql/db"
	"github.com/idprm/go-payment/src/datasource/redis"
	"github.com/idprm/go-payment/src/domain/repository"
	"github.com/idprm/go-payment/src/handler"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/services"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Worker CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * Load config
		 */
		cfg, err := config.LoadSecret("secret.yaml")
		if err != nil {
			panic(err)
		}

		/**
		 * Init DB
		 */
		db, err := db.InitMySQL(cfg)
		if err != nil {
			panic(err)
		}

		/**
		 * Init redis
		 */
		rds, err := redis.InitRedis("")
		if err != nil {
			panic(err)
		}

		/**
		 * Init log
		 */
		lg := logger.NewLogger(cfg)
		zap := logger.InitLogger(cfg)

		ctx := context.Background()

		// init order
		orderRepo := repository.NewOrderRepository(db)
		orderService := services.NewOrderService(orderRepo)

		// init transaction
		transactionRepo := repository.NewTransactionRepository(db)
		transactionService := services.NewTransactionService(transactionRepo)

		// init callback
		callbackRepo := repository.NewCallbackRepository(db)
		callbackService := services.NewCallbackService(callbackRepo)
		handler.NewCallbackHandler(cfg, rds, lg, zap, orderService, transactionService, callbackService)

		fmt.Println("Ready message")

		for {
			result, err := rds.BRPop(ctx, 0, Q_PAY).Result()
			if err != nil {
				fmt.Println(err)
				continue
			}

			// var req *entity.HandlerRequest
			// json.Unmarshal([]byte(result[1]), &req)

			// print
			fmt.Println(result[1])

			// h.Send(req)

			// Wait a random amount of time before popping the next item
			time.Sleep(time.Duration(rand.Intn(8)) * time.Second)
		}
	},
}