package migrator

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate(table interface{}, db *gorm.DB) error {
	if err := db.AutoMigrate(&table); err != nil {
		fmt.Printf("error when migrating %v", err)
		return err
	}

	return nil
}
