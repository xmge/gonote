package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go MemoryLeak()
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func MemoryLeak()  {
	i := 1
	for {
		r, _ := http.Get("xmge.top")
		_ = r
		fmt.Println(i)
		//r.Body.Close()//
		time.Sleep(time.Second * 1)
		i++
	}
}
