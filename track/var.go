package main

import (
	"fmt"
	"math"
)

var I = 1
// a := 1 // 'a' unexpected
func init()  {
	fmt.Println(I)
}

func main() {

	// var  声明赋值变量
	// var 可以在方法外使用
	var a int
	var b int = 2 	//Type can be omitted
	var c string = "a" //Type can be omitted
	var d,e = 5,"b"


	// := 声明变量
	// := 不可在方法外使用，并且左边赋值的变量必须有未被申明过的变量，否则无法使用
	// a,b := 1,2 // No new variables on left side of :=
	a,f := 1,2

	fmt.Println(a,b,c,d,e,f)

	// 冒号声明的易错点
	//
	trackDemo()
}

func trackDemo()  {
	var datas = []float64{1,3,4,2}
	max := datas[0]
	for _, data := range datas[1:] {
		max := math.Max(max,data)
		_=max
	}
	fmt.Println(max)
}
