package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// 几种对称性加密算法：AES,DES,3DES
// 几种非对称性加密算法：RSA,DSA,ECC
// 几种线性散列算法（签名算法）：MD5,SHA1,HMAC

func main() {
	str := `{"user_id":718721389}s7D8scdf89Dcs1dL0vrfFk02fSkd02fC`
	Md5(str)
	Sha1(str)
}

func Md5(str string)  {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(str))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}

func Sha1(str string) {
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(str))
	Result := Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}

