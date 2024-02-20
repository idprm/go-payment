package cmd

import (
	"log"
	"os"

	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

const (
	APP_NAME                string = "APP_NAME"
	APP_URL                 string = "APP_URL"
	APP_HOST                string = "APP_HOST"
	APP_PORT                string = "APP_PORT"
	APP_TZ                  string = "APP_TZ"
	URI_PGSQL               string = "URI_PGSQL"
	URI_MYSQL               string = "URI_MYSQL"
	URI_REDIS               string = "URI_REDIS"
	DRAGONPAY_URL           string = "DRAGONPAY_URL"
	DRAGONPAY_MERCHANTID    string = "DRAGONPAY_MERCHANTID"
	DRAGONPAY_PASSWORD      string = "DRAGONPAY_PASSWORD"
	DRAGONPAY_POSTBACK      string = "DRAGONPAY_POSTBACK"
	JAZZCASH_URL            string = "JAZZCASH_URL"
	JAZZCASH_MERCHANTID     string = "JAZZCASH_MERCHANTID"
	JAZZCASH_PASSWORD       string = "JAZZCASH_PASSWORD"
	JAZZCASH_INTEGERITYSALT string = "JAZZCASH_INTEGERITYSALT"
	MIDTRANS_URL            string = "MIDTRANS_URL"
	MIDTRANS_MERCHANTID     string = "MIDTRANS_MERCHANTID"
	MIDTRANS_CLIENTKEY      string = "MIDTRANS_CLIENTKEY"
	MIDTRANS_SERVERKEY      string = "MIDTRANS_SERVERKEY"
	MOMO_URL                string = "MOMO_URL"
	MOMO_PARTNERCODE        string = "MOMO_PARTNERCODE"
	MOMO_ACCESSKEY          string = "MOMO_ACCESSKEY"
	MOMO_SECRETKEY          string = "MOMO_SECRETKEY"
	NICEPAY_URL             string = "NICEPAY_URL"
	NICEPAY_MERCHANTID      string = "NICEPAY_MERCHANTID"
	NICEPAY_MERCHANTKEY     string = "NICEPAY_MERCHANTKEY"
	RAZER_URL               string = "RAZER_URL"
	RAZER_MERCHANTID        string = "RAZER_MERCHANTID"
	RAZER_VERIFYKEY         string = "RAZER_VERIFYKEY"
	RAZER_SECRETKEY         string = "RAZER_SECRETKEY"
	XIMPAY_URL_TSEL         string = "XIMPAY_URL_TSEL"
	XIMPAY_URL_H3I          string = "XIMPAY_URL_H3I"
	XIMPAY_URL_XL           string = "XIMPAY_URL_XL"
	XIMPAY_URL_ISAT         string = "XIMPAY_URL_ISAT"
	XIMPAY_URL_SF           string = "XIMPAY_URL_SF"
	XIMPAY_URL_XLPIN        string = "XIMPAY_URL_XLPIN"
	XIMPAY_URL_SFPIN        string = "XIMPAY_URL_SFPIN"
	XIMPAY_PARTNERID        string = "XIMPAY_PARTNERID"
	XIMPAY_SECRETKEY        string = "XIMPAY_SECRETKEY"
	XIMPAY_USERNAME         string = "XIMPAY_USERNAME"
	LOG_PATH                string = "LOG_PATH"
)

const (
	Q_PAY string = "q_payment"
)

func init() {
	/**
	 * WEBSERVER SERVICE
	 */
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(listenerCmd)

	/**
	 * WORKER SERVICE
	 */
	rootCmd.AddCommand(workerCmd)
	rootCmd.AddCommand(seederCmd)

	rootCmd.AddCommand(testCmd)

}

func Execute() error {
	return rootCmd.Execute()
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panicf("Error %v", key)
	}
	return value
}

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(getEnv(URI_MYSQL)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// DEBUG ON CONSOLE
	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate(
		&entity.Country{},
		&entity.Gateway{},
		&entity.Credential{},
		&entity.Application{},
		&entity.Channel{},
		&entity.Order{},
		&entity.Payment{},
		&entity.Callback{},
		&entity.Refund{},
		&entity.Transaction{},
		&entity.Return{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectRedis() (*redis.Client, error) {
	opts, err := redis.ParseURL(getEnv(URI_REDIS))
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opts), nil
}
