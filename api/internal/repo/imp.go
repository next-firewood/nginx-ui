package repo

import (
	"context"
	"gorm.io/gorm"
)

type GlobalRepo interface {
	BeginTx(ctx context.Context) (context.Context, func(err error) error)
	GetDB(ctx context.Context) *gorm.DB

	SysConfDber
	UserDber
}

const _txKey = "db_tx"

func (r *globalDb) BeginTx(ctx context.Context) (context.Context, func(err error) error) {
	tx := r.DB.Begin()

	ctx = context.WithValue(ctx, _txKey, tx)

	clean := func(err error) error {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

		return err
	}

	return ctx, clean
}

func (r *globalDb) GetDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(_txKey).(*gorm.DB); ok && tx != nil {
		if tx.Statement == nil || tx.Statement.ConnPool == nil {
			return r.DB.WithContext(ctx)
		}

		return tx.WithContext(ctx)
	}

	return r.DB.WithContext(ctx)
}
