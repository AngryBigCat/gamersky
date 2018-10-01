package models

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/jinzhu/gorm"
)

const DATATYPE = "mysql"
const DATAHOST = "localhost"
const DATANAME = "gamersky"
const DATAROOT = "root"
const DATAPASS = "123456"
const DATAPORT = 3306

var DB *gorm.DB

func init() {
	conn, err := gorm.Open(DATATYPE,
		fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
			DATAROOT,
			DATAPASS,
			DATAHOST,
			DATAPORT,
			DATANAME,
		))
	checkErr(err)
	DB = conn
}

func checkErr(err error) {
	if err != nil {
		log.Error("database error: %s", err)
	}
}