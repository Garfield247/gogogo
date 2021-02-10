package main

import "fmt"

func main() {
	m1 := Mouse{name: "logic G304"}
	testInterFace(m1)
	f1 := FlashDisk{name: "SanDisk 64Gb"}
	testInterFace(f1)
}

// 接口
type USB interface {
	start() // 开始工作
	end()   // 结束工作
}

// 实现类
type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Printf("鼠标:%s已连接\n", m.name)
}

func (m Mouse) end() {
	fmt.Printf("鼠标:%s已断开连接\n", m.name)
}

func (f FlashDisk) start() {
	fmt.Printf("U盘:%s已插入\n", f.name)
}

func (f FlashDisk) end() {
	fmt.Printf("U盘:%s已安全移除\n", f.name)
}

func testInterFace(usb USB) {
	usb.start()
	usb.end()
}
