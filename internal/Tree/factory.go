package Tree

import (
	"fmt"
	"strings"
)

func NewIntTree() *Tree {
	return &Tree{Compare: IntCompare}
}

func IntCompare(d1 interface{}, d2 interface{}) int {
	if d1.(int) > d2.(int) {
		return 1
	}
	if d1.(int) == d2.(int) {
		return 0
	}
	return -1
}

func NewStringTree() *Tree {
	return &Tree{Compare: StringCompare}
}

func StringCompare(d1 interface{}, d2 interface{}) int {
	return strings.Compare(fmt.Sprintf("%v", d1), fmt.Sprintf("%v", d2))
}
