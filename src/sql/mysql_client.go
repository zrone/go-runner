package sql

import (
	"awesome-runner/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"sync"
)

var (
	MDB *gorm.DB
)

// initialize mysql
// singleton
func initMysqlClient() *gorm.DB {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Cnf.MysqlDNS + "&time_zone=" + url.QueryEscape(config.Cnf.TimeZone), // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                                                                        // add default size for string fields, by default, will use db types `longtext` for fields without size, not transform primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                       // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                       // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                       // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                      // smart configure based on used version
	}), &gorm.Config{
		PrepareStmt: true,
	})

	return db
}

// get *gorm.DB
func GetMysqlInstance() *gorm.DB {
	if MDB == nil {
		var once sync.Once
		once.Do(func() {
			MDB = initMysqlClient()
		})
	}
	return MDB
}
