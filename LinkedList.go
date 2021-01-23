package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func NewNode(Val int) *Node {
	t := Node{Val, nil}
	return &t
}

func (Head *Node) Push(Val int) *Node {
	if Head == nil {
		Head = NewNode(Val)
		return Head
	}
	Helper := Head
	for Head.Next != nil {
		Head = Head.Next
	}
	Head.Next = NewNode(Val)
	Head.Next = Head.Next
	return Helper
}

func (Head *Node) Pop() (*Node, bool) {
	if Head == nil {
		return nil, false
	}
	Head = Head.Next
	return Head, true
}

func (Head *Node) Print() {
	for Head != nil {
		fmt.Print(Head.Value, " ")
		Head = Head.Next
	}
}

func AssertEqual(lhs bool, rhs bool, msg string) {
	if lhs != rhs {
		fmt.Println("Testcase error! Line: " + msg)
	}
}

func testAll() {
	var Head *Node
	var ok bool
	Head = Head.Push(5)
	Head = Head.Push(15)
	Head = Head.Push(13)
	Head = Head.Push(14)
	fmt.Println("Before pop: ")
	Head.Print()
	Head, ok = Head.Pop()
	fmt.Println("\nAfter pop 1: ")
	Head.Print()
	AssertEqual(ok, true, "61")
	Head, ok = Head.Pop()
	fmt.Println("\nAfter pop 2: ")
	Head.Print()
	AssertEqual(ok, true, "65")
	Head, ok = Head.Pop()
	fmt.Println("\nAfter pop 3: ")
	Head.Print()
	AssertEqual(ok, true, "69")
	Head, ok = Head.Pop()
	fmt.Println("\nAfter pop 4: ")
	Head.Print()
	AssertEqual(ok, true, "73")
	Head, ok = Head.Pop()
	fmt.Println("\nAfter fake pop 5: ")
	Head.Print()
	AssertEqual(ok, false, "77")
}

func main() {
	//cin := bufio.NewReader(os.Stdin)
	testAll()

}
