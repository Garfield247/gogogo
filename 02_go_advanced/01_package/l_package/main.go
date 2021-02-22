// Package main provides ...
package main

import (
	"fmt"
	"l_package/person"
	_ "l_package/pkg1" // 使用空白标识符可以只导包初始化且不调用 的情况下不报错
	"l_package/utils"
	tu "l_package/utils/timeutils" // 给包起别名 类似 Python  import as
)

func main() {
	utils.Count()
	utils.Test()
	tu.EchoTime()

	p1 := person.Person{Name: "小明", Sex: "nan"}
	fmt.Println(p1)
	p1.Info()
}
