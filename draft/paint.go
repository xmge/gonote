package main

import (
	"fmt"
	"math/rand"
	"time"
)


const (
	up = iota
	down
	left
	right
)

var _directions = []int{up,down,left,right}

func main() {
	fmt.Println(GenMap(100))
}


func GenMap(pointNum int)  map[int]int {

	defer func() {
		if err := recover();err != nil {
			fmt.Println()
		}
	}()
	coordinates := make(map[int]int)

	coordinates[0]=0
	for i := 0; i < pointNum; i++ {
		var x,y int
		directions := make([]int,4)
		copy(directions,_directions)
		ok := true
		for ok {
			if len(directions) == 0{
				return coordinates
			}
			rand.Seed(time.Now().Unix())
			randNum := rand.Intn(len(directions))
			direction := directions[randNum]
			switch direction {
			case up:
				x+=1
			case down:
				x-=1
			case left:
				y-=1
			case right:
				y+=1
			}
			y,ok = coordinates[x]
			if !ok {
				fmt.Println(x,y)
				coordinates[x]=y
			}else {
				directions=append(directions[:randNum],directions[randNum+1:]...)
			}
		}
	}
	return coordinates
}
