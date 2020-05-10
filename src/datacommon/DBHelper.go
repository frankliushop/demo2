package datacommon

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//開啟DB
func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:frank99@tcp(192.168.56.101:3306)/personDB?parseTime=true&loc=Local")

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
