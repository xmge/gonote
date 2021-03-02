package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type Num struct {
		i string
		j int64
		m bool
	}
	n := Num{"ABC", 123, true}
	nPointer := unsafe.Pointer(&n)

	fmt.Println("i的偏移量：", unsafe.Offsetof(n.i))

	fmt.Println("i占几个字节：", unsafe.Sizeof(n.i))
	fmt.Println("j的偏移量：", unsafe.Offsetof(n.j))
	fmt.Println("j占几个字节：", unsafe.Sizeof(n.j))
	fmt.Println("m的偏移量：", unsafe.Offsetof(n.m))
	fmt.Println("m占几个字节：", unsafe.Sizeof(n.m))

	niPointer := (*string)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.i)))
	*niPointer = "QWE"
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 456
	nmPointer := (*bool)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.m)))
	*nmPointer = false

	fmt.Println(n.i)
	fmt.Println(n.j)
	fmt.Println(n.m)
}


func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

