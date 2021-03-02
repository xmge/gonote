
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

	return ret(i)
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

	return ret(i)
}

func ret(i int) int {
	fmt.Println("call return function")
	i++
	return i
}


// test1:
//  2
// 2020/11/19 17:38:55 this is panic
// test2:
//  0
// 2020/11/19 17:38:55 this is panic
// test3:
//  1
```