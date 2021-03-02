package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"sync"
)

type User struct {
	Name string `validate:"required"`
	Age int `validate:"gt=0"`
}

func main() {
	validate := validator.New()
	wg := sync.WaitGroup{}
	for i := 0; i <100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j <10; j++ {
				user:= User{Name: "xmge",Age: 0}
				fmt.Println(i,j,validate.Struct(user))
			}
		}(i)
	}
	wg.Wait()
}
