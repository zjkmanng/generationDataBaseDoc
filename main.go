package main

import (
	"encoding/json"
	"fmt"
	"generationDataBaseDoc/db"
	"generationDataBaseDoc/do"
)

func main() {

	db.InitEngine()
	//metas, err := db.Engine.DBMetas()
	//if err != nil {
	//	return
	//}
	////JSON序列化：结构体-->JSON格式的字符串
	//data, err := json.Marshal(metas)
	//if err != nil {
	//	fmt.Println("json marshal failed")
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("json:%s\n", data)
	var user do.TUser
	table, err := db.Engine.TableInfo(&user)
	if err != nil {
		return
	}
	data, err := json.Marshal(table)
	if err != nil {
		fmt.Println("json marshal failed")
		fmt.Println(err)
		return
	}
	fmt.Printf("json:%s\n", data)
}
