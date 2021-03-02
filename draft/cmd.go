package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 通过exec.Command函数执行命令或者shell
	// 第一个参数是命令路径，当然如果PATH路径可以搜索到命令，可以不用输入完整的路径
	// 第二到第N个参数是命令的参数
	// 下面语句等价于执行命令: ls -l /var/
	cmd := exec.Command("ls", "-l", "/var/")
	// 执行命令，并返回结果
	output,err := cmd.Output()
	if err != nil {
		panic(err)
	}
	// 因为结果是字节数组，需要转换成string
	fmt.Println(string(output))

}