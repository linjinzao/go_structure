package mysql

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

var db *sql.DB

func init(){
	db,_ = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	
	db.SetMaxOpenConns(1000)
	
	err := db.Ping()
	if err != nil {
		fmt.Println("连接失败：err:",err)
		os.Exit(1)
	}
}

func DBConn() *sql.DB{
	return db
}