package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	runtime.Caller()
}

