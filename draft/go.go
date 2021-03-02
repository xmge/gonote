package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		go func() {
			i := 1
			for  {
				time.Sleep(1*time.Second)
				i++
				fmt.Println(i)
			}
		}()
		fmt.Println("over")
	}()


	select {}
}
