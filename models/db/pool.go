package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/silenceper/pool"
)

var Pool pool.Pool

func InitPool() {

	//factory 创建连接的方法
	factory := func() (interface{}, error) {
		return gorm.Open(DATATYPE,
			fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8",
				DATAROOT,
				DATAPASS,
				DATAHOST,
				DATAPORT,
				DATANAME,
			))
	}

	//close 关闭连接的方法
	close := func(db interface{}) error {
		return db.(*gorm.DB).Close()
	}

	poolConfig := &pool.Config{
		InitialCap: 30,
		MaxCap:     100,
		Factory:    factory,
		Close:      close,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 5 * time.Second,
	}

	p, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("err=", err)
	}

	Pool = p
}
