package main

import (
	"fmt"
)

type Human interface {
	Eat()
}

type Student struct {}

func (s Student)Eat(){} // 值方法

type Teacher struct {}

func (t *Teacher)Eat(){} // 指针方法

func main() {
	var a Human = &Student{}
	var b Human = Student{}
	var c Human = &Teacher{}
	//var d Human = Teacher{}  无法赋值
	fmt.Println(a,b,c)
}

// TODO 为什么会出现这样的情况