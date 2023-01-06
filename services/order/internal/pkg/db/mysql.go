package db

import (
	"context"
	"github.com/h-varmazyar/snappfood/pkg/errors"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/entity"
	"google.golang.org/grpc/codes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMySql(ctx context.Context, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("can not load repository configs")
	}

	if err = migrateModels(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(
		new(entity.Order),
	); err != nil {
		return err
	}
	return nil
}
