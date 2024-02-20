package cmd

import (
	"context"
	"log"

	"github.com/idprm/go-payment/src/app"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/datasource/mysql/db"
	"github.com/idprm/go-payment/src/datasource/redis"
	"github.com/idprm/go-payment/src/logger"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Webserver CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * LOAD CONFIG
		 */
		cfg, err := config.LoadSecret("secret.yaml")
		if err != nil {
			panic(err)
		}

		log.Println(cfg)

		/**
		 * Init DB
		 */
		db, err := db.InitMySQL(cfg)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Init redis
		rds, err := redis.InitRedis(cfg)
		if err != nil {
			panic(err)
		}

		/**
		 * Init Log
		 */
		lg := logger.NewLogger(cfg)
		zap := logger.InitLogger(cfg)

		ctx := context.Background()

		application := app.NewApplication(cfg, db, rds, lg, zap, ctx)
		router := application.Start()
		log.Fatal(router.Listen(":" + cfg.App.Port))
	},
}
