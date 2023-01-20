package Tree

import (
	"io"
)

type Tree struct {
	Root    *Node
	Compare func(d1 interface{}, d2 interface{}) int
}

func (t *Tree) Insert(data interface{}) *Tree {
	if t.Root == nil {
		t.Root = &Node{Data: data, Right: nil, Left: nil, Height: 0}
	} else {
		t.Root.insert(data, 0, t.Compare)
	}

	return t
}

func (t *Tree) Find(data interface{}) int {
	if t.Compare(t.Root.Data, data) == 0 {
		return t.Root.Height
	}

	return t.Root.find(data, t.Compare)
}

func (t *Tree) Print(w io.Writer) {
	t.Root.print(w, "ROOT")
}
