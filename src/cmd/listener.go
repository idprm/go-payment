package cmd

import (
	"context"
	"log"

	"github.com/idprm/go-payment/src/app"
	"github.com/idprm/go-payment/src/logger"
	"github.com/spf13/cobra"
)

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Webserver CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * Init DB
		 */
		db, err := connectDB()
		if err != nil {
			panic(err)
		}

		// Init redis
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * Init Log
		 */
		lg := logger.NewLogger()
		zap := logger.InitLogger()

		ctx := context.Background()

		application := app.NewApplication(db, rds, lg, zap, ctx)
		router := application.Start()
		log.Fatal(router.Listen(":" + APP_PORT))
	},
}
