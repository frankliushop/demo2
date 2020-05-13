package datacommon

import (
	"database/sql"
	"log"
	"myproject1/dataconfig"

	_ "github.com/go-sql-driver/mysql"
)

var DbCon *sql.DB

//初始化方法
func init() {
	var err error
	DbCon, err = sql.Open("mysql", dataconfig.GlobalConfigData.SqlConnection)
	if err != nil {
		log.Fatalln(err)
	}

	DbCon.SetMaxIdleConns(200)
	DbCon.SetMaxOpenConns(100)

	if err = DbCon.Ping(); err != nil {
		log.Fatalln(err)
	}
}

//計算總頁數
func GetPageCount(pageSize int64,itemCount int64) int64 {
	pageCount := itemCount / pageSize
	if itemCount % pageSize > 0{
		pageCount++
	}
	return pageCount
}
