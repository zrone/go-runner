package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

type Manager struct {
	ID       uint64 `json:"id"`
	Role     uint8  `json:"role"`
	Username string `json:"username"`
	Account  string `json:"account"`
	Salt     string `json:"salt"`
	CreateAt uint32 `json:"create_at"`
	DeleteAt uint32 `json:"delete_at"`
	UpdateAt uint32 `json:"update_at"`
}

func (m Manager) TableName() string {
	return "manager"
}

func main() {
	timezone := "'Asia/Shanghai'"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(localhost:3306)/pave?charset=utf8mb4&parseTime=True&loc=Local&time_zone=" + url.QueryEscape(timezone), // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                                                                                                                    // add default size for string fields, by default, will use db types `longtext` for fields without size, not transform primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                                                                   // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                                   // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                                   // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                                  // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Manager{})

	// Create
	//db.Create(&Manager{
	//	Role:     1,
	//	Username: "admin",
	//	Account:  "张三",
	//	Salt:     "demo@salt",
	//	CreateAt: uint32(time.Now().Unix()),
	//	DeleteAt: 0,
	//	UpdateAt: uint32(time.Now().Unix()),
	//})

	// Read
	//var manager Manager
	//db.First(&manager, 1)                       // 根据整形主键查找
	//db.First(&manager, "username = ?", "admin") // 查找 username 字段值为 admin 的记录

	// Update - 将 product 的 price 更新为 200
	//db.Model(&manager).Update("Account", "李四")
	//// Update - 更新多个字段
	//db.Model(&manager).Updates(Manager{Salt: "pose@salt"}) // 仅更新非零值字段
	//db.Model(&manager).Updates(map[string]interface{}{"Salt": "salt@salt"})

	//// Delete - 删除 product
	//db.Delete(&manager, 1)

	fmt.Println("hello world!!!")
}
