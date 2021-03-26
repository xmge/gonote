# defer 可以修改 return 的返回值吗？

## 问题答案

如果被调用的函数，返回值有声明变量名，则 defer 可以修改返回值，否则无法修改。

因为如果函数未声明返回值变量名，那么函数在 return 时，会把返回结果保存在一个临时变量 temp 中，但是在 defer 时却
拿不到这个变量，因此 defer 无法修改

但是如果函数中声明了返回值变量名，那么函数在 return 时，会把返回值保存在这个变量中，那么 defer 就可以通过修改这个
变量来修改返回值了

## 示例 Demo

```go
package main

import (
	"errors"
	"fmt"
	"log"
)

// return 和 defer panic 关于返回值详解

// return：是把一个变量保存到零时变量，如果函数返回值中申明了变量，这会将结果保存到这个变量中。
// defer：是在 return 后执行的，如果函数返回值中没有申明变量，则 defer 中无法修改返回值，如果函数返回值中有申明变量，则 defer 可以通过修改这个变量来修改返回值
// panic: 如果 panic 则不会执行 return 语句，直接执行 defer 语句，执行 defer 语句，如果 defer 语句没有对变量进行改变，则默认返回返回类型的零值。

func main() {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
}

// 函数返回值没有申明变量，
// return 时，把 i 保存进了 temp 变量（临时保存变量的地方）
// defer 时，把 i 进行了+1，由 2 变为了 3，但是并没有改变 temp 的值，所有结果还是 2
func test1() int {
	var i = 1
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		i++
	}()

	i++
	return i
}

// 函数返回值没有申明变量，
// return 之前有 panic,所以没有执行 return 语句，因此返回值中为 temp=0 （返回类型的默认值）
// defer 时，把 i 进行了+1，由 0 变为了 1，但是并没有改变 temp 的值，所有结果还是 0
func test2() int {
	var i = 1
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		i++
	}()

	panic(errors.New("this is panic"))

	return i
}

// 函数返回值申明变量了变量 i
// return 之前有 panic,所以没有执行 return 语句，因此返回值中为 i=0 （返回类型的默认值）
// defer 时，把 i 进行了+1，由 0 变为了 1，此时 i 就是保存返回结果的变零，因此返回结果为 1
func test3() (i int) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		i++
	}()

	panic(errors.New("this is panic"))

	return i
}

//
//2
//2021/03/11 22:35:48 this is panic
//0
//2021/03/11 22:35:48 this is panic
//1


```
