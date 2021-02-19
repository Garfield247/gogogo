package main

import "fmt"

func main() {
	//length, width := 3.2, 4.7
	//length, width := -3.2, 4.7
	//length, width := 3.2, -4.7
	length, width := -3.2, -4.7
	area, err := ciraArea(length, width)
	if err != nil {
		//fmt.Println(err)
		if instance, ok := err.(*areaError); ok {
			if instance.lengthNegative() {
				fmt.Printf("高:%.2f不合法\n", instance.length)
			}
			if instance.widthNegative() {
				fmt.Printf("宽:%.2f不合法\n", instance.width)
			}
		}
		return
	}
	fmt.Println(area)
}

type areaError struct {
	msg    string
	length float64
	width  float64
}

func (e *areaError) Error() string {
	return e.msg
}
func (e *areaError) lengthNegative() bool {
	return e.length < 0
}
func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func ciraArea(length, width float64) (float64, error) {
	msg := ""
	if length < 0 {
		msg = "高不合法"
	}
	if width < 0 {
		if msg == "" {
			msg = "宽不合法"
		} else {
			msg += ",宽不合法"
		}
	}
	if msg != "" {
		return 0, &areaError{msg: msg, length: length, width: width}
	}
	return length * width, nil

}
