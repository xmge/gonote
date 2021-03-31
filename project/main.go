package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func main() {
	A()
}

func A()  {
	err := errors.New("aaa")
	HandlerError(err)
}
func HandlerError(err error)  {
	fmt.Println(string(debug.Stack()))
}
