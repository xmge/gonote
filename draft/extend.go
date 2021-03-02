package main

import (
	"fmt"
	"time"
)

type BaseModel struct {
	Created time.Time `json:"created"`
	Updated time.Time  `json:"updated"`
}

func (m * BaseModel)Create()  {
	fmt.Println(m)
}

type UserModel struct {
	BaseModel
	Username string
	Age int
}

func main() {
	u := UserModel{Username: "xmge",Age: 1}
	u.Create()
}
