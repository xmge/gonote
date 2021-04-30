package language

import "fmt"

// go语言各个类型是否可比较大全
func main() {
	// 数字类型比较
	var i1,i2 = 1,2
	fmt.Println(i1==i2)

	// bool 类型比较
	var b1,b2 = true,false
	fmt.Println(b1==b2)

	// string 类型比较
	var str1,str2= "a","b"
	fmt.Println(str1==str2)

	// 同类型数组类型比较
	var array1,array2 = [1]int{},[1]int{}
	fmt.Println(array1==array2)

	// 数组类型 x
	var array3,array4 = [1]int{},[2]int{}
	//Invalid operation: array3==array4 (mismatched types [1]int and [2]int)
	fmt.Println(array3==array4)

	// 切片类型比较 x
	var slice1,slice2 = []int{},[]int{}
	// Invalid operation: slice1==slice2 (operator == is not defined on []int)
	fmt.Println(slice1==slice2)

	// map 比较 x
	var m1,m2 = map[string]interface{}{},map[string]interface{}{}
	// Invalid operation: m1==m2 (operator == is not defined on map[string]interface{})
	fmt.Println(m1==m2)

	// channel 比较
	var c1,c2 = make(chan bool,2),make(chan bool)
	fmt.Println(c1==c2)

	// function 比较 x
	var f1,f2 =F(),F()
	//Invalid operation: f1 == f2 (operator == is not defined on func())
	fmt.Println(f1 == f2)

	// interface 比较
	var if1,if2 interface{}
	fmt.Println(if1==if2)

	// struct 比较
	// struct 有些是可以比较的，有些不可以比较
	// 可比较的结构体
	// 不包含 slice,map,function 的结构可以比较，在比较时会依次比较每个字段的值
	// 包含 slice,map,function 的结构提不可比较
}

type F func()