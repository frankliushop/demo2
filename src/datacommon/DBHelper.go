package datacommon

import (
	"database/sql"
	"myproject1/dataconfig"

	_ "github.com/go-sql-driver/mysql"
)

//開啟DB
func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", dataconfig.GlobalConfigData.SqlConnection)
	if err != nil {
		panic(
			ExceptionData{
				ErrorCode:    ErrCodeDBCanNotOpen,
				ErrorMessage: err.Error(),
			},
		)
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		//log.Fatalln(err)
		panic(
			ExceptionData{
				ErrorCode:    ErrCodeDBCanNotOpen,
				ErrorMessage: err.Error(),
			},
		)
	}
	return db
}
