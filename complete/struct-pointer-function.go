package main

import "fmt"

//
type Person struct {
	Name string
	Age int
	Sex bool
}

func (p *Person)PointerChangeName()  {
	p.Name="change1"
}

func (p Person)ObjectChangeName2()  {
	p.Name="change2"
}

func main() {

	// 指针方法
	p1 := Person{}
	p1.PointerChangeName()
	fmt.Println(p1)

	p2 := &Person{}
	p2.PointerChangeName()
	fmt.Println(p2)

	// 对象方法
	p3 := Person{}
	p3.ObjectChangeName2()
	fmt.Println(p3)

	p4 := Person{}
	p4.ObjectChangeName2()
	fmt.Println(p4)
}


//{change1 0 false}
//&{change1 0 false}
//{ 0 false}
//{ 0 false}

//结论：
//1. 无论是指针方法还是对象方法，对应结构体的指针和对象均可以对它们进行调用。
//2. 并且会在调用的过程中，自动转为调用方法接受类型，比如 object 类型调用指针方法时，会把 &object 传入方法中，在 &object 类型调用对象方法时，会把 object 传进去。
//3. 因此，调用指针方法传进去的就是对象的指针，调用对象方法传进去的就是对象，与调用者是怎样的类型没有关系，go 会自动做转换
