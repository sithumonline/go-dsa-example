package tree

import "fmt"

type Node struct {
	index int
	data  float32
	max   *Node
	min   *Node
}

func NewNode(index int, data float32) *Node {
	return &Node{
		index: index,
		data:  data,
	}
}

func (n *Node) display() {
	fmt.Printf("%d: %f\n", n.index, n.data)
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(index int, data float32) {
	n := NewNode(index, data)
	if t.root == nil {
		t.root = n
		return
	}
	current := t.root
	for {
		if index > current.index {
			if current.max != nil {
				current = current.max
				continue
			}
			current.max = n
			return
		}

		if index < current.index {
			if current.min != nil {
				current = current.min
				continue
			}
			current.min = n
			return
		}

		return
	}
}

func (t *Tree) Search(index int) *Node {
	current := t.root
	for {
		if current == nil {
			return nil
		}

		if index > current.index {
			current = current.max
			continue
		}

		if index < current.index {
			current = current.min
			continue
		}

		return current
	}
}

func (t *Tree) Delete(index int) {
	current := t.root
	for {
		if current == nil {
			return
		}

		if index > current.index {
			current = current.max
			continue
		}

		if index < current.index {
			current = current.min
			continue
		}

		if current.min == nil {
			current.max = nil
			return
		}

		if current.max == nil {
			current.min = nil
			return
		}

		current.min.max = current.max
		current.max.min = current.min
		current.min = nil
		current.max = nil
		return
	}
}

func (t *Tree) Display(index int) {
	current := t.root
	for {
		if current == nil {
			return
		}

		if index > current.index {
			current = current.max
			continue
		}

		if index < current.index {
			current = current.min
			continue
		}

		current.display()
		return
	}
}
