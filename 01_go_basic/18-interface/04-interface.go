// Package main provides ...
package main

import (
	"fmt"
	"math"
)

// 接口断言
func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Printf("周长:%f,面积:%f\n", t1.peri(), t1.area())
	var c1 Circle = Circle{5}
	fmt.Printf("周长:%f,面积:%f\n", c1.peri(), c1.area())
	var s1 Shape
	s1 = t1
	fmt.Printf("周长:%f,面积:%f\n", s1.peri(), s1.area())
	var s2 Shape
	s2 = c1
	fmt.Printf("周长:%f,面积:%f\n", s2.peri(), s2.area())

	testShape(s1)
	testShape(s2)
	testShape(t1)
	testShape(c1)

	getType(s1)
	getType(s2)
	getType(c1)
	getType(t1)

	var t2 *Triangle = &Triangle{3, 4, 5}
	var c2 *Circle = &Circle{4}
	getType(t2)
	getType(c2)

	getTypeBySwitch(s1)
	getTypeBySwitch(s2)
	getTypeBySwitch(c1)
	getTypeBySwitch(t1)
	getTypeBySwitch(t2)
	getTypeBySwitch(c2)
}

func getType(s Shape) {
	if instance, ok := s.(Triangle); ok {
		fmt.Printf("三角形%v\n", instance)
		fmt.Printf("三角形%T\n", instance)
	} else if instance, ok := s.(Circle); ok {
		fmt.Printf("圆形%v\n", instance)
	} else if instance, ok := s.(*Triangle); ok {
		fmt.Printf("三角形的指针类型|%v|%T\n", instance, instance)
	} else if instance, ok := s.(*Circle); ok {
		fmt.Printf("圆形的指针类型|%v|%T\n", instance, instance)
	} else {
		fmt.Println("未知")
	}
}

func getTypeBySwitch(s Shape) {
	switch instance := s.(type) {
	case Triangle:
		fmt.Println("三角形")
		fmt.Printf("%T|%v|%p\n", instance, instance, &instance)
	case Circle:
		fmt.Println("圆形")
		fmt.Printf("%T|%v|%p\n", instance, instance, &instance)
	case *Triangle:
		fmt.Println("三角形结构体指针")
		fmt.Printf("%T|%v|%p\n", instance, instance, &instance)
	case *Circle:
		fmt.Println("圆形的指针类型")
		fmt.Printf("%T|%v|%p\n", instance, instance, &instance)
	default:
		fmt.Println("未知")
		fmt.Printf("%T|%v|%p", instance, instance, &instance)

	}
}

func testShape(s Shape) {
	fmt.Printf("周长:%f,面积:%f\n", s.peri(), s.area())

}

type Shape interface {
	peri() float64
	area() float64
}

type Triangle struct {
	a, b, c float64
}

type Circle struct {
	radius float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p * t.c))
}

func (c Circle) peri() float64 {
	return 2 * c.radius * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
