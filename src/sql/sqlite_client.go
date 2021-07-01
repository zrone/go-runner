package sql

import (
	"awesome-runner/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var SDB *gorm.DB

// initialize mysql
// singleton
func initLiteClient() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("data/pave.db"), &gorm.Config{
		PrepareStmt: true,
	})

	db.AutoMigrate(&types.InternalDeploy{})
	return db
}

// get *gorm.DB
func GetLiteInstance() *gorm.DB {
	if SDB == nil {
		var once sync.Once
		once.Do(func() {
			SDB = initLiteClient()
		})
	}
	return SDB
}
