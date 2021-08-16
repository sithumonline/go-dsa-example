package linklist

import "fmt"

type List struct {
	data int
	next *List
}

func newList(n int) *List {
	return &List{
		data: n,
	}
}

func (l *List) display() {
	fmt.Printf("%d\n", l.data)
}

type LinkList struct {
	first *List
}

func NewLinkList() *LinkList {
	return &LinkList{}
}

func (l *LinkList) InsertFirst(n int) {
	current := newList(n)
	current.next = l.first
	l.first = current
}

func (l *LinkList) DeleteFirst() {
	current := l.first
	l.first = current.next
}

func (l *LinkList) IsEmpty() bool {
	return l.first == nil
}

func (l *LinkList) DisplayAll() {
	current := l.first
	for current.next != nil {
		current.display()
		current = current.next
	}
}

func (l *LinkList) InsertMid(id int, n int) {
	current := l.first
	for current.next != nil {
		if current.data == id {
			tmp := current
			l := newList(n)
			l.next = tmp.next
			current.next = l
			break
		}
		current = current.next
	}
}
