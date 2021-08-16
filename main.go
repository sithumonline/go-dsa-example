package main

import (
	"fmt"

	"github.com/sithumonline/go-dsa-example/linkList"
	"github.com/sithumonline/go-dsa-example/queue"
	"github.com/sithumonline/go-dsa-example/stack"
	"github.com/sithumonline/go-dsa-example/tree"
)

func main() {
	s := stack.NewStack(3)
	s.Push(1)
	fmt.Printf("%v\n", s.Peek())
	s.Push(2)
	s.Pop()
	fmt.Printf("%v\n", s.IsEmpty())
	fmt.Printf("%v\n", s.IsFull())
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Printf("%v\n", s.IsFull())
	fmt.Printf("%v\n", s.Peek())
	s.Pop()
	fmt.Printf("%v\n", s.Peek())

	q := queue.NewQueue(3)
	q.Insert(145)
	q.Insert(666)
	q.Insert(1234)

	fmt.Printf("%v\n", q.Remove())
	fmt.Printf("%v\n", q.PeekFront())
	fmt.Printf("%v\n", q.Remove())
	fmt.Printf("%v\n", q.PeekFront())

	list := linklist.NewLinkList()
	fmt.Println(list.IsEmpty())
	list.InsertFirst(100)
	list.InsertFirst(200)
	fmt.Println(list.IsEmpty())
	list.InsertFirst(300)
	list.DisplayAll()

	t := tree.NewTree()
	t.Insert(100, 78)
	t.Insert(200, 18)
	t.Insert(300, 92)
	t.Insert(400, 12)
	t.Display(200)
	t.Insert(200, 12)
	t.Display(200)
}
