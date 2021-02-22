// Package mydb provides ...
package mydb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() {

	mydqldb, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mydqldb)
}
