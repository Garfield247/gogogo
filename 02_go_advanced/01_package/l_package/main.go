// Package main provides ...
package main

import (
	"fmt"
	"l_package/person"
	"l_package/utils"
	tu "l_package/utils/timeutils" // 给包起别名 类似 Python  import as
)

func main() {
	utils.Count()
	tu.EchoTime()

	p1 := person.Person{Name: "小明", Sex: "nan"}
	fmt.Println(p1)
	p1.Info()
}
