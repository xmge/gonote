# pprof 排查内存泄漏问题

## 1. 如何发现内存泄漏?

内存泄漏是随着服务运行时间变长，在运行时有些资源没有释放，导致占用内存越来越大的现象

因此如果有内存监控的话可以更容易发现内存泄漏的问题，这里推荐大家使用 tick （telegraf-influxdb-grafana）
来进行系统的监控，如果有内存泄漏，监控显示大概是这样的(绿色的线为内存，每次重启内存都会释放)：

![](http://qpmy9rk30.hb-bkt.clouddn.com/tick_memory_leak.png)

## 2. 模拟内存泄漏

```go
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
func MemoryLeak()  {
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
```

## 3. 通过 pprof 排查内存泄漏的原因

排查步骤：

1. 点击 run 方法，让服务跑起来
2. 通过命令获取当前内存情况 `curl localhost:8080/debug/pprof/heap > heap.base`
3. 等待一段时间，10s 中左右，通过命令获取目前内存情况 `curl localhost:8080/debug/pprof/heap > heap.current`
4. 通过对比两次内存的情况，来查看两次内存情况哪里增长比较多，找到内存泄漏的地方 `pprof -http=:8000 -base heap.base heap.current`
5. 如图：
![](http://qpmy9rk30.hb-bkt.clouddn.com/pprof_memory_leak.png)
6. 可以发现 ppprof 工具已经提示是 MemoryLeak ，泄漏的地方就是 ioutil.ReadAll() 方法导致的，因此我们去找到相应的代码进行修改就可以修复问题了。