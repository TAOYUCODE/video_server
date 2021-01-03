package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:060017@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
}