package Tree

import (
	"fmt"
	"io"
	"strings"
)

type Node struct {
	Data   interface{}
	Right  *Node
	Left   *Node
	Height int
}

func (n *Node) insert(data interface{}, h int, compare func(d1 interface{}, d2 interface{}) int) {
	h++
	if data == nil {
		return
	} else {
		switch compare(data, n.Data) {
		case 1:
			if n.Left == nil {
				n.Left = &Node{Data: data, Left: nil, Right: nil, Height: h}
			} else {
				n.Left.insert(data, h, compare)
			}
		case -1:
			if n.Right == nil {
				n.Right = &Node{Data: data, Left: nil, Right: nil, Height: h}
			} else {
				n.Right.insert(data, h, compare)
			}
		}
	}
}

func (n *Node) find(data interface{}, compare func(d1 interface{}, d2 interface{}) int) int {
	switch compare(data, n.Data) {
	case 0:
		return n.Height
	case 1:
		if n.Left == nil {
			return -1
		}
		return n.Left.find(data, compare)
	case -1:
		if n.Right == nil {
			return -1
		}
		return n.Right.find(data, compare)
	}

	return -1
}

func (n *Node) print(w io.Writer, prefix string) {
	if n == nil {
		return
	}
	fmt.Fprintf(w, "%s: %v [H:%d]\n", prefix, n.Data, n.Height)
	n.Left.print(w, fmt.Sprintf(" %sL", strings.Repeat(" ", n.Height)))
	n.Right.print(w, fmt.Sprintf(" %sR", strings.Repeat(" ", n.Height)))
}
