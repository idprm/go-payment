package cmd

import (
	"log"
	"os"

	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
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

var (
	APP_NAME  string = getEnv("APP_NAME")
	APP_URL   string = getEnv("APP_URL")
	APP_HOST  string = getEnv("APP_HOST")
	APP_PORT  string = getEnv("APP_PORT")
	APP_TZ    string = getEnv("APP_TZ")
	URI_PGSQL string = getEnv("URI_PGSQL")
	URI_MYSQL string = getEnv("URI_MYSQL")
	URI_REDIS string = getEnv("URI_REDIS")
	LOG_PATH  string = getEnv("LOG_PATH")
)

const (
	Q_PAY string = "q_payment"
)

func init() {
	/**
	 * WEBSERVER SERVICE
	 */
	rootCmd.AddCommand(listenerCmd)

	/**
	 * WORKER SERVICE
	 */
	rootCmd.AddCommand(consumerCmd)
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
	db, err := gorm.Open(mysql.Open(URI_MYSQL), &gorm.Config{})
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
	opts, err := redis.ParseURL(URI_REDIS)
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opts), nil
}

func connectPgSQL() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(URI_PGSQL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// DEBUG ON CONSOLE
	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate(
		&entity.Application{},
		&entity.Gateway{},
		&entity.Channel{},
		&entity.Order{},
		&entity.Transaction{},
		&entity.Payment{},
		&entity.Callback{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil

}
