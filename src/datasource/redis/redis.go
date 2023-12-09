package redis

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Secret) (*redis.Client, error) {
	opts, err := redis.ParseURL(cfg.Redis.Url)
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opts), nil
}
