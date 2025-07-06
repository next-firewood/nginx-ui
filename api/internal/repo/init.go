package repo

import (
	"gorm.io/gorm"
	"log"
	"rui/common/database"
	"sync"
)

var dbs *gorm.DB
var once sync.Once

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&SysConf{},
	)
}

func GetDB(conf *database.DBConf) *gorm.DB {
	once.Do(func() {
		db, err := conf.Init()
		if err != nil {
			log.Fatalln(err)
		}

		dbs = db
	})

	return dbs
}

type globalDb struct {
	DB *gorm.DB
}

func NewGlobalDb(DB *gorm.DB) GlobalRepo {
	return &globalDb{
		DB: DB,
	}
}
