package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go	TestCostTime()
	port := "8080"
	fmt.Printf("Now listening  on: http://0.0.0.0:%s\n",port)
	log.Fatal(http.ListenAndServe(":"+port,nil))
}

func TestCostTime()  {
	tick := time.Tick(time.Millisecond * 100)
	for range tick {
		ShortTime()
		LongTime()
	}
}

func ShortTime()  {
	time.Sleep(time.Millisecond)
}

func LongTime()  {
	time.Sleep(time.Millisecond * 100)
}