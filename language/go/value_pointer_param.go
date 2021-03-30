package main

import (
	"fmt"
	"reflect"
)


//对于Go语言，严格意义上来讲，只有一种传递，也就是按值传递(by value)。当一个变量当作参数传递的时候，会创建一个变量的副本，然后传递给函数或者方法，你可以看到这个副本的地址和变量的地址是不一样的。

func main() {

	// map 指针传递 map[string]int 和 *map[string]int 是有区别的，不过在传递参数时会
	m := map[string]int{"a":1}
	fmt.Printf("原始map的内存地址是：%p\n", m)
	fmt.Printf("原始map的内存地址是：%p\n", &m)
	ChangeMap(m)
	fmt.Println(m)

	// channel 指针传递
	c := make(chan int,1)
	ChangeChannel(c)
	fmt.Println(len(c))

	// slice 指针传递
	s := make([]int,0)
	s = append(s, 1,2)
	ChangeSlice(s)
	fmt.Println(s)

	// array 值传递
	a := [2]int{1,2}
	ChanceArray(a)
	fmt.Println(a)


}

func ChangeMap(m map[string]int){
	fmt.Printf("原始map的内存地址是：%p\n", m)
	fmt.Printf("原始map的内存地址是：%p\n", &m)
	fmt.Println(reflect.TypeOf(m))
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

func ChangeChannel(c chan int)  {
	fmt.Println(reflect.TypeOf(c))
	c<-1
}

//map[a:1 b:2]
//[]int
//[100 2]
//[2]int
//[1 2]
