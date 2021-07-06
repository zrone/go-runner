package sql

import (
	"awesome-runner/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"sync"
	"time"
)

var SDB *gorm.DB

// initialize mysql
// singleton
func initLiteClient() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("data/pave.db"), &gorm.Config{
		PrepareStmt: true,
	})
	db.Use(dbresolver.Register(dbresolver.Config{ /* xxx */ }).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(100).
		SetMaxOpenConns(200))

	db.AutoMigrate(&types.InternalDeploy{}, &types.TaskLog{})
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
