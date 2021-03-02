package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func main()  {
	n1,n2,n3,n4,n5  := new(Node),new(Node),new(Node),new(Node),new(Node)
	n1.Value=1
	n2.Value=2
	n3.Value=3
	n4.Value=4
	n5.Value=5
	n1.Next=n2
	n2.Next=n3
	n3.Next=n4
	n4.Next=n5
	list(n1)
	findMiddleNode(n1)
}


// 链表遍历
func list(n *Node)  {
	for n.Next != nil {
		fmt.Println(n.Value)
		n = n.Next
	}
}

// 链表找中间值
func findMiddleNode(n *Node)  {
	slowN,fastN := n,n
	for fastN.Next != nil{
		slowN=slowN.Next
		fastN=fastN.Next.Next
	}
	fmt.Println(slowN)
}