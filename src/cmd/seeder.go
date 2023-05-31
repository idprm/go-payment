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

		var country []entity.Country
		var application []entity.Application
		var gateway []entity.Gateway
		var channel []entity.Channel

		var countries = []entity.Country{
			{Name: "Indonesia", Locale: "id", Prefix: "62", Flag: "indonesia.png"},
			{Name: "Philippine", Locale: "ph", Prefix: "63", Flag: "philippine.png"},
			{Name: "Pakistan", Locale: "pk", Prefix: "92", Flag: "pakistan.png"},
			{Name: "Vietnam", Locale: "vn", Prefix: "84", Flag: "vietnam.png"},
			{Name: "Malaysia", Locale: "my", Prefix: "60", Flag: "malaysia.png"},
			{Name: "Thailand", Locale: "th", Prefix: "66", Flag: "thailand.png"},
		}

		var applications = []entity.Application{
			{
				CountryID:   1,
				Code:        "sehatcepat",
				Name:        "Sehat Cepat",
				Url:         "https://www.sehatcepat.com",
				UrlCallback: "https://www.sehatcepat.com/payment/callback",
				UrlReturn:   "https://www.sehatcepat.com/order",
			},
			{
				CountryID:   2,
				Code:        "pagalingkaagad",
				Name:        "Pagalingkaagad",
				Url:         "https://www.pagalingkaagad.com",
				UrlCallback: "https://www.pagalingkaagad.com/payment/callback",
				UrlReturn:   "https://www.pagalingkaagad.com/order",
			},
			{
				CountryID:   2,
				Code:        "pagalingkaagad",
				Name:        "Pagalingkaagad",
				Url:         "https://www.pagalingkaagad.com",
				UrlCallback: "https://www.pagalingkaagad.com/payment/callback",
				UrlReturn:   "https://www.pagalingkaagad.com/order",
			},
			{
				CountryID:   1,
				Code:        "SURATSAKIT",
				Name:        "Surat Sakit",
				Url:         "https://www.suratsakit.com",
				UrlCallback: "https://www.suratsakit.com/payment/callback",
				UrlReturn:   "https://www.suratsakit.com/order",
			},
			{
				CountryID:   2,
				Code:        "GEMEZZVN",
				Name:        "CP GEMEZZ",
				Url:         "https://vngemezz.exmp.app",
				UrlCallback: "https://vngemezz.exmp.app",
				UrlReturn:   "https://vngemezz.exmp.app",
			},
		}

		var gateways = []entity.Gateway{
			{CountryID: 1, Code: "MIDTRANS", Name: "Midtrans", Currency: "IDR"},
			{CountryID: 1, Code: "NICEPAY", Name: "Nicepay", Currency: "IDR"},
			{CountryID: 2, Code: "DRAGONPAY", Name: "Dragon Pay", Currency: "PHP"},
			{CountryID: 3, Code: "JAZZCASH", Name: "Jazz Cash", Currency: "PKR"},
			{CountryID: 4, Code: "MOMO", Name: "Momo Payment", Currency: "VND"},
			{CountryID: 5, Code: "RAZER", Name: "Razer", Currency: "MYR"},
		}

		var channels = []entity.Channel{
			{GatewayID: 1, Name: "GoPay", Slug: "gopay", Logo: "gopay.png", Type: "wallet", Param: "gopay", IsActive: true},
			{GatewayID: 1, Name: "ShopeePay", Slug: "shopeepay", Logo: "shopeepay.png", Type: "wallet", Param: "shopeepay", IsActive: true},
			{GatewayID: 1, Name: "BNI Virtual Account", Slug: "bni-va", Logo: "bni-va.png", Type: "wallet", Param: "bank-transfer/bni-va", IsActive: true},
			{GatewayID: 1, Name: "BCA Virtual Account", Slug: "bca-va", Logo: "bca-va.png", Type: "wallet", Param: "bank-transfer/bca-va", IsActive: true},
			{GatewayID: 1, Name: "BRI Virtual Account", Slug: "bri-va", Logo: "bri-va.png", Type: "wallet", Param: "bank-transfer/bri-va", IsActive: true},
			{GatewayID: 1, Name: "Mandiri", Slug: "mandiri-bill", Logo: "mandiri-bill.png", Type: "wallet", Param: "bank-transfer/mandiri-bill", IsActive: true},
			{GatewayID: 1, Name: "Permata Virtual Account", Slug: "permata-va", Logo: "permata-va.png", Type: "bank-transfer/permata-va", Param: "PYMY", IsActive: true},
			{GatewayID: 2, Name: "OVO", Slug: "ovo", Logo: "ovo.png", Type: "wallet", Param: "OVOE", IsActive: true},
			{GatewayID: 2, Name: "DANA", Slug: "dana", Logo: "dana.png", Type: "wallet", Param: "DANA", IsActive: true},
			{GatewayID: 2, Name: "LinkAja", Slug: "linkaja", Logo: "linkaja.png", Type: "wallet", Param: "LINK", IsActive: true},
			{GatewayID: 3, Name: "BDO Internet Banking", Slug: "bdo", Logo: "bdologo.jpg", Type: "transfer", Param: "BDO", IsActive: true},
			{GatewayID: 3, Name: "Bogus Bank", Slug: "bogus", Logo: "boguslogo.jpg", Type: "transfer", Param: "BOG", IsActive: true},
			{GatewayID: 3, Name: "BPI Online/Mobile", Slug: "bpi", Logo: "bpilogo.jpg", Type: "transfer", Param: "BPIA", IsActive: true},
			{GatewayID: 3, Name: "GCash App", Slug: "gcash", Logo: "gcashlogo.jpg", Type: "wallet", Param: "GCSH", IsActive: true},
			{GatewayID: 3, Name: "Metrobankdirect", Slug: "metrod", Logo: "metrodlogo.jpg", Type: "transfer", Param: "MBTC", IsActive: true},
			{GatewayID: 3, Name: "BogusBank OTC", Slug: "bogus", Logo: "boguslogo.jpg", Type: "transfer", Param: "BOGX", IsActive: true},
			{GatewayID: 3, Name: "Bank of Commerce Online", Slug: "boc", Logo: "boclogo.jpg", Type: "transfer", Param: "BOC", IsActive: true},
			{GatewayID: 3, Name: "WeChat Pay", Slug: "wechat", Logo: "wechatlogo.jpg", Type: "wallet", Param: "AUWC", IsActive: true},
			{GatewayID: 3, Name: "EastWest Online/Cash Payment", Slug: "ewblogo", Logo: "ewblogo.jpg", Type: "transfer", Param: "EWXB", IsActive: true},
			{GatewayID: 3, Name: "GrabPay", Slug: "grabpay", Logo: "grabpaylogo.jpg", Type: "wallet", Param: "GRPY", IsActive: true},
			{GatewayID: 3, Name: "PayMaya", Slug: "paymaya", Logo: "paymayalogo.jpg", Type: "wallet", Param: "PYMY", IsActive: true},
			{GatewayID: 4, Name: "Jazz Cash", Slug: "jazzcash", Logo: "jazzcash.png", Type: "wallet", Param: "JCASH", IsActive: true},
			{GatewayID: 5, Name: "Momo Wallet", Slug: "momo-wallet", Logo: "momo.png", Type: "wallet", Param: "MOMO", IsActive: true},
			{GatewayID: 6, Name: "Visa MasterCard", Slug: "visa-master", Logo: "visa-master.png", Type: "wallet", Param: "index.php", IsActive: true},
			{GatewayID: 6, Name: "Bank Islam", Slug: "bank-islam", Logo: "bank-islam.png", Type: "transfer", Param: "BIMB.php", IsActive: true},
			{GatewayID: 6, Name: "Public Bank", Slug: "public-bank", Logo: "public-bank.png", Type: "transfer", Param: "PBB.php", IsActive: true},
			{GatewayID: 6, Name: "Maybank2u", Slug: "maybank", Logo: "maybank.png", Type: "transfer", Param: "maybank2u.php", IsActive: true},
			{GatewayID: 6, Name: "Hong Leong", Slug: "hong-leong", Logo: "hong-leong.png", Type: "transfer", Param: "hlb.php", IsActive: true},
			{GatewayID: 6, Name: "CIMB Clicks", Slug: "cimb-clicks", Logo: "cimb-clicks.png", Type: "transfer", Param: "cimb.php", IsActive: true},
			{GatewayID: 6, Name: "RHB Now", Slug: "rhb-now", Logo: "rhb-now.png", Type: "wallet", Param: "rhb.php", IsActive: true},
			{GatewayID: 6, Name: "7 Eleven", Slug: "7-eleven", Logo: "7-eleven.png", Type: "other", Param: "cash.php", IsActive: true},
			{GatewayID: 6, Name: "AmOnline", Slug: "am-online", Logo: "am-online.png", Type: "wallet", Param: "amb.php", IsActive: true},
			{GatewayID: 6, Name: "Razer Pay", Slug: "razer-pay", Logo: "razer-pay.png", Type: "wallet", Param: "RazerPay.php", IsActive: true},
			{GatewayID: 6, Name: "Affin Bank", Slug: "affin-bank", Logo: "affin-bank.png", Type: "transfer", Param: "affin-epg.php", IsActive: true},
		}

		if db.Find(&country).RowsAffected == 0 {
			for i, _ := range countries {
				db.Model(&entity.Country{}).Create(&countries[i])
			}
			log.Println("countries migrated")
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
