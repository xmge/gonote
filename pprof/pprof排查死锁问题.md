# pprof 排查死锁问题

## 1. 死锁的四个必要条件

1. 互斥条件
2. 请求与保持
3. 不剥夺条件
4. 循环与等待

## 2. 如何发现死锁？

死锁会使请求阻塞，程序无法继续执行下去，因此最明显的现象就是服务正常运行但是服务请求超时，因此当出现大量请求超时的时候
要重点考虑下是不是死锁造成的

## 3. 程序模拟死锁现象

```go
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
```

## 4. 通过 pprof 排查死锁问题原因

在查询 Deadlock 时，主要看 goroutine 是不是飙升，如果程序中 goroutine 的数量和并发量有较大的出入，可以直接通过浏览器访问 /debug/pprof 来进行查看当前 goroutine 的数量，如图所示：

![goroutine-飙升](http://qiniu.gonote.cn/goroutine-%E9%A3%99%E5%8D%87.png)

可以发现 goroutine 数量达到了 8502个，而且还在不断增加，因此怀疑是死锁问题，点击进去可以查看 goroutine 当前的分布，如图所示：

![goroutine-详情](http://qiniu.gonote.cn/goroutine-detail-7.09.png)

可以看出， 在 `/Users/maning/docs/draft/deadlock.go:37` 和 
`/Users/maning/docs/draft/deadlock.go:43` 各有 827 个 goroutine 在等待，因此我们要到这两个地方进行调查原因，
查看为什么 goroutine 在这里阻塞，从而找到死锁的问题代码，然后解决问题。