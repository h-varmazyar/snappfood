package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DB struct {
	MySqlDB *gorm.DB
	RedisDB *redis.Client
}

func NewDatabase(ctx context.Context, configs *Configs) (*DB, error) {
	db := new(DB)
	if configs.MySqlDSN != "" {
		postgres, err := newMySql(ctx, configs.MySqlDSN)
		if err != nil {
			return nil, err
		}
		db.MySqlDB = postgres
	}

	if configs.RedisAddress != "" {
		db.RedisDB = redis.NewClient(&redis.Options{
			Addr:     configs.RedisAddress,
			Password: configs.RedisPassword,
			DB:       configs.RedisLockDB,
		})
	}

	return db, nil
}
