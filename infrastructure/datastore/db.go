package datastore

import (
	"fmt"
	"log"

	"github.com/GustavoGutierrez/goapi/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	strConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%s&loc=Local", config.C.Database.User, config.C.Database.Password, config.C.Database.Net, config.C.Database.Addr, config.C.Database.DBName, config.C.Database.Params.ParseTime)

	// fmt.Println(strConn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       strConn, // data source name
		DefaultStringSize:         256,     // default size for string fields
		DisableDatetimePrecision:  true,    // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,    // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,    // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,   // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln("ERROR DB:", err)
	}

	return db
}
