package main

import (
	"fmt"
	"reflect"
)

func i() {
	type I interface {}
	type User struct {}

	var i I
	fmt.Println(i==nil)
	var u *User
	i = u
	fmt.Println(u==nil)	// 值为空就为 nil
	fmt.Println(i==nil) // 类型和值都为空才为 nil
	fmt.Println(reflect.TypeOf(i))
}

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

func main() {
	var w worker = person{}
	fmt.Println(w)
}
