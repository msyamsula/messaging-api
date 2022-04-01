package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var Db *gorm.DB

type PgConfig struct {
	Host     string
	Port     string
	Dbname   string
	User     string
	Password string
}

func Connect(config PgConfig) error {

	uri := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.Host,
		config.User,
		config.Password,
		config.Dbname,
		config.Port,
	)

	dbConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy:         nil,
		FullSaveAssociations:   false,
		Logger:                 nil,
		NowFunc: func() time.Time {
			return time.Now()
		},
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           map[string]clause.ClauseBuilder{},
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  map[string]gorm.Plugin{},
	}

	db, err := gorm.Open(postgres.Open(uri), dbConfig)
	if err != nil {
		panic("panic")
	}
	Db = db
	return nil
}

func GetDB() *gorm.DB {
	return Db
}
