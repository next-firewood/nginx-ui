package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"time"
)

type DBConf struct {
	DSN             string
	ConnMaxIdleTime time.Duration `json:",default=10m"`
	MaxIdleConns    int           `json:",default=20"`  // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	MaxOpenConns    int           `json:",default=20"`  // SetMaxOpenConns 设置打开数据库连接的最大数量。
	ConnMaxLifetime time.Duration `json:",default=10m"` // SetConnMaxLifetime 设置了连接可复用的最大时间。
}

// Init 初始化数据库连接
func (c *DBConf) Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(c.DSN), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		FullSaveAssociations:   true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	if err = db.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxIdleTime(c.ConnMaxIdleTime).
			SetConnMaxLifetime(c.ConnMaxLifetime).
			SetMaxIdleConns(c.MaxIdleConns).
			SetMaxOpenConns(c.MaxOpenConns),
	); err != nil {
		return nil, err
	}

	return db, nil
}
