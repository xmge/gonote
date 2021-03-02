package tools

import "syscall"

const (
	lf = "singleton.lock"
)

// 保证应用是单例的
func Singleton() {
	if fd, err := syscall.Open(lf, syscall.O_CREAT|syscall.O_RDONLY, 0644); err != nil {
		panic(err)
	} else if err := syscall.Flock(fd, syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		panic(err)
	}
	return
}
