package main

import (
	"fmt"
	"math"
)

func main() {
	m := make(map[float64]int)
	m[1.6] = 1 // 10行
	m[math.NaN()] = 2 // 11行
	m[2]=2
	fmt.Println(m[1.6])
	fmt.Println(m[1.60000001])
	fmt.Println(m[1.60000000000000001])
	fmt.Println(m[math.NaN()])

	_,ok := m[1.6]
	fmt.Println(ok)
	_,ok = m[1123]
	fmt.Println(ok)
}

