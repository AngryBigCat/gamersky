package db

import (
	"fmt"
	"log"

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

func Get() *gorm.DB {
	db, err := gorm.Open(DATATYPE,
		fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8",
			DATAROOT,
			DATAPASS,
			DATAHOST,
			DATAPORT,
			DATANAME,
		))

	if err != nil {
		log.Fatalln("数据库连接失败")
	}

	DB = db
	return db
}
