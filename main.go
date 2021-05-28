package main

import (
	"fmt"
	"generationDataBaseDoc/db"
)

func main() {
	// 开启数据库连接
	db.InitDB()
	tableMap := make(map[string]string)
	tableMap = TableInfo("zhangjie")
	fmt.Println(tableMap)
}

func TableInfo(dbName string) map[string]string {
	sqlStr := `SELECT table_name tableName,TABLE_COMMENT tableDesc
			FROM INFORMATION_SCHEMA.TABLES 
			WHERE UPPER(table_type)='BASE TABLE'
			AND LOWER(table_schema) = ? 
			ORDER BY table_name asc`

	var result = make(map[string]string)

	rows, err := db.Db.Query(sqlStr, dbName)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var tableName, tableDesc string
		err = rows.Scan(&tableName, &tableDesc)
		if err != nil {
			fmt.Println(err)
		}

		if len(tableDesc) == 0 {
			tableDesc = tableName
		}
		result[tableName] = tableDesc
	}
	return result
}
