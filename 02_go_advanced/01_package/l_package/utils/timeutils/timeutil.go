// Package timeutil provides ...
package timeutil

import (
	. "fmt" // 使用点使用时可以省略保命
	"time"
)

func EchoTime() {
	Println(time.Now())
}
