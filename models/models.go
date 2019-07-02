package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const DATATYPE = "mysql"
const DATAHOST = "111.230.26.237"
const DATANAME = "gamersky"
const DATAROOT = "root"
const DATAPASS = "g8LSfVZ53chK4cTY"
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
	}
}
