package tools

import (
	"os"
	"os/signal"
	"syscall"
)

// 监听信号
func WatchSignal(callback func(sig os.Signal),sig ...os.Signal) {
	sc := make(chan os.Signal, 3)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR2,
	)

	go func() {
		for {
			sig := <-sc
			handleSignal(sig)
		}
	}()
}

func handleSignal(sig os.Signal) {
	if sig == syscall.SIGUSR2 {
		// 重新打开日志文件，比如用 logrotate 压缩日志后，需要程序重新创建日志文件
	} else {
		os.Exit(0)
	}
}
