package cmd

import (
	"log"

	"github.com/idprm/go-payment/src/app"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/datasource/pgsql/db"
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
		db, err := db.InitDB(cfg)
		if err != nil {
			log.Fatal(err)
			return
		}

		/**
		 * Init Log
		 */
		// logger := logger.InitLogger(conf)

		application := app.NewApplication(cfg, db)
		router := application.Start()
		log.Fatal(router.Listen(":" + cfg.App.Port))
	},
}
