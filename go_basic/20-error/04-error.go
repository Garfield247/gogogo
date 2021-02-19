package main

import (
	"fmt"
	"math"
)

func main() {
	radiu := -3.0
	area, err := circArea(radiu)
	if err != nil {
		//fmt.Println(err)
		if instance, ok := err.(*areaError); ok {
			fmt.Printf("msg:%s,radiu:%.2f", instance.msg, instance.radiu)
		} else {
			fmt.Println(err)
		}
		return
	}
	fmt.Printf("面积:%.2f", area)

}

type areaError struct {
	radiu float64
	msg   string
}

func (e *areaError) Error() string {
	return fmt.Sprintf("error:radiu:%.2f,%s", e.radiu, e.msg)
}

func circArea(radiu float64) (float64, error) {
	if radiu < 0 {
		return 0, &areaError{msg: "半径不合法", radiu: radiu}
	}
	return math.Pow(radiu, 2) * math.Pi, nil
}
