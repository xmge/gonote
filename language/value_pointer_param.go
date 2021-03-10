package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]int{"a":1}
	ChangeMap(m)
	fmt.Println(m)

	s := make([]int,0)
	s = append(s, 1,2)
	ChangeSlice(s)
	fmt.Println(s)

	a := [2]int{1,2}
	ChanceArray(a)
	fmt.Println(a)
}

func ChangeMap(m map[string]int){
	m["b"]=2
}

func ChangeSlice(s []int){
	fmt.Println(reflect.TypeOf(s))
	s[0]=100
}

func ChanceArray(a [2]int)  {
	fmt.Println(reflect.TypeOf(a))
	a[0]=100
}

//map[a:1 b:2]
//[]int
//[100 2]
//[2]int
//[1 2]
