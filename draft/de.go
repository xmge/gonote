package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//BHOCGPSHBZI
	fmt.Println(BHex2Num("BHOCGPSHBZI",36))
	fmt.Println(strconv.ParseInt("BHOCGPSHBZI",36,64))
}


var num2char = "0123456789abcdefghijklmnopqrstuvwxyz"


// 10进制数转换   n 表示进制， 16 or 36
func NumToBHex(num, n int) string {
	num_str := ""
	for num != 0 {
		yu := num % n
		num_str = string(num2char[yu]) + num_str
		num = num / n
	}
	return strings.ToUpper(num_str)
}

// 36进制数转换   n 表示进制， 16 or 36
func BHex2Num(str string, n int) int {
	str = strings.ToLower(str)
	v := 0.0
	length := len(str)
	for i := 0; i < length; i++ {
		s := string(str[i])
		index := strings.Index(num2char, s)
		v += float64(index) * math.Pow(float64(n), float64(length-1-i)) // 倒序
	}
	return int(v)
}
