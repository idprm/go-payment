package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/redis/go-redis/v9"
)

type VerfifyRepository struct {
	rds *redis.Client
}

func NewVerfifyRepository(rds *redis.Client) *VerfifyRepository {
	return &VerfifyRepository{
		rds: rds,
	}
}

type IVerifyRepository interface {
	Get(string) (*entity.Verify, error)
	Set(*entity.Verify) error
	Del(*entity.Verify) error
}

func (r *VerfifyRepository) Get(key string) (*entity.Verify, error) {
	val, err := r.rds.Get(context.TODO(), key).Result()
	if err != nil {
		return nil, err
	}
	var v *entity.Verify
	json.Unmarshal([]byte(val), &v)
	return v, nil
}

func (r *VerfifyRepository) Set(v *entity.Verify) error {
	verify, _ := json.Marshal(v)
	err := r.rds.Set(context.TODO(), v.GetKey(), string(verify), 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *VerfifyRepository) Del(v *entity.Verify) error {
	err := r.rds.Del(context.TODO(), v.GetKey()).Err()
	if err != nil {
		return err
	}
	return nil
}
