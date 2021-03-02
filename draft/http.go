package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Print("\r",i)
	}
}