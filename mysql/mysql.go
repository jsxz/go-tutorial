package main

import (
	"database/sql"
	"fmt"

	//只调用这个包的init方法
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, e := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	if e != nil {
		panic(e)
	}
	stmt, e1 := db.Prepare("insert user_info set username=?,departname=?,create_time=?")
	if e1 != nil {
		panic(e1)
	}
	result, _ := stmt.Exec("zhang", "技术部", "2019-11-11")
	id, _ := result.LastInsertId()
	r1, _ := result.RowsAffected()
	fmt.Println(id)
	fmt.Println(r1)

	rows, _ := db.Query("select * from user_info")
	for rows.Next() {
		var id int
		var username string
		var departname string
		var create_time string
		_ = rows.Scan(&id, &username, &departname, &create_time)
		fmt.Println(id, username, departname, create_time)
	}

}
