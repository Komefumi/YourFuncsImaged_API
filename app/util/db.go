package util

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDBAccess() func() *gorm.DB {
	dsn :=  os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASS") + "@/" + os.Getenv("MYSQL_DB") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic (err)
	}

	return func() *gorm.DB {
		fmt.Println(dsn)
		return db
	}
}

// DBAccessorFunc to access the GORM database context
var DBAccessorFunc = getDBAccess()