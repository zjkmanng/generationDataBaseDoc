package db

import (
	"fmt"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitEngine() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("数据库链接错误日志：")
		fmt.Println("/n")
		fmt.Println(err)
	}
}
