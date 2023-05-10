package cmd

import (
	"log"

	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/datasource/mysql/db"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "Seeder CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * LOAD CONFIG
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
			log.Fatal(err)
			return
		}

		var application []entity.Application
		var gateway []entity.Gateway
		var channel []entity.Channel

		var applications = []entity.Application{
			{Code: "SEHATCEPAT", Name: "Sehat Cepat", Domain: "sehatcepat.com", UrlCallback: ""},
			{Code: "SURATSAKIT", Name: "Surat Sakit", Domain: "suratsakit.com", UrlCallback: ""},
		}

		var gateways = []entity.Gateway{
			{Code: "DRAGONPAY", Name: "Dragon Pay"},
			{Code: "JAZZCASH", Name: "Jazz Cash"},
			{Code: "MIDTRANS", Name: "Midtrans"},
			{Code: "MOMO", Name: "Momo Payment"},
			{Code: "NICEPAY", Name: "Nicepay"},
			{Code: "RAZER", Name: "Razer"},
		}

		var channels = []entity.Channel{
			{Name: "", Param: "", IsActive: true},
		}

		if db.Find(&application).RowsAffected == 0 {
			for i, _ := range applications {
				db.Model(&entity.Application{}).Create(&applications[i])
			}
			log.Println("applications migrated")
		}

		if db.Find(&gateway).RowsAffected == 0 {
			for i, _ := range gateways {
				db.Model(&entity.Gateway{}).Create(&gateways[i])
			}
			log.Println("gateways migrated")
		}

		if db.Find(&channel).RowsAffected == 0 {
			for i, _ := range channels {
				db.Model(&entity.Channel{}).Create(&channels[i])
			}
			log.Println("channels migrated")
		}

	},
}
