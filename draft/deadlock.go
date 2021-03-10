package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)



func main() {
	go SendRequest()
	http.HandleFunc("/",DeadLineHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func SendRequest()  {
	tick := time.Tick(time.Millisecond * 100)
	for range tick {
		_,err := http.Get("http://localhost:8080")
		if err != nil {
			log.Panic(err)
		}
	}
}

func DeadLineHandler(writer http.ResponseWriter, request *http.Request) {
	tick := time.Tick(time.Millisecond * 10)
	for range tick {
		var m1,m2 sync.Mutex
		c := make(chan bool)
		go func() {
			m1.Lock()
			c<-true
			m2.Lock()
		}()

		go func() {
			m2.Lock()
			<-c
			m1.Lock()
		}()
	}
}
