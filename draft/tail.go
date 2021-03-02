package main

import (
	"fmt"
	"github.com/nxadm/tail"
	"os"
	"time"
)

func main() {

	go func() {
		f,err := os.Create("app.log")
		if err != nil {
			panic(err)
		}

		for {
			f.WriteString("ERROR , db error\n")
			f.WriteString("this is info log\n")
			f.WriteString("this is debug log\n")
			time.Sleep(time.Second*1)
		}
	}()

	t, err := tail.TailFile("app.log", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Println(line)
		//if strings.Contains(line.Text,"ERROR") {
		//	fmt.Println("ERROR",line.Text)
		//}
	}
}

