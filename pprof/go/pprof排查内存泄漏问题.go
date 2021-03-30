
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go MemoryLeak()
	log.Fatal(http.ListenAndServe(":8080",nil))
}

// 通过多个 goroutine 发送很多 http 请求，并且不调用 r.body.close() 来进行资源释放，从而导致内存泄漏
// 实验方法：
// 1. 点击 run 方法，让服务跑起来
// 2. 通过命令获取当前内存情况 curl localhost:8080/debug/pprof/heap > heap.base
// 3. 等待一段时间，10s 中左右，通过命令获取目前内存情况 curl localhost:8080/debug/pprof/heap > heap.current
// 4. pprof -http=:8000 -base heap.base heap.current 进行内存比对，来查看两次内存情况哪里增长比较多，找到内存泄漏的地方
// 5. 如图：
func MemoryLeak ()  {
	// 知识点： 为什么不调用 r.Body.Close() 就会导致内存泄漏？
	// 因为 http 请求时启用了 3 个goroutine,主G，writeG，readG，如果不 close() 掉的话，
	// readG 会阻塞直到收到 EOF,而调用 close() 方法 就是让 readG 结束，从而释放资源

	tick := time.Tick(time.Millisecond *100)
	for range tick{
		go func() {
			r ,err := http.Get("http://www.baidu.com")
			if err != nil {
				log.Panic(err)
			}
			//defer r.Body.Close() // 没有释放资源，导致内存泄漏
			_,err = ioutil.ReadAll(r.Body)
			if err != nil{
				log.Panic(err)
			}
		}()
	}
}

