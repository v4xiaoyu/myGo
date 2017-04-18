package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func CloseDb() {
	db.Close()
}

func ConnectDb() {
	db, err = sql.Open("mysql", "root:942698.Liu@/mytestdb")
	//数据库连接字符串，别告诉我看不懂。端口一定要写/
	if err != nil {
		//连接成功 err一定是nil否则就是报错
		//	panic(err.Error()) //抛出异常
		fmt.Println(err.Error()) //仅仅是显示异常
	}
	//defer CloseDb()  //只有在前面用了 panic 这时defer才能起作用，如果链接数据的时候出问题，他会往err写数据
}

func GetStmt(sqlStr string) *sql.Stmt {
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err.Error()) //仅仅是显示异常
	}

	return stmt
}
