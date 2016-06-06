package trees

import "fmt"

type Node interface {
	fmt.Stringer
	Value() interface{}
	Left() Node
	Right() Node
}

type Comparator func(v1, v2 interface{}) int

type InsertFunc func(*Tree, interface{})

type RemoveFunc func(*Tree, interface{}) bool

type Tree struct {
	Root Node
	C    Comparator
	i InsertFunc
	r RemoveFunc
}

func (t *Tree) String() string {
	if t.Root == nil {
		return "EMPTY TREE"
	}
	return t.Root.String()
}

func (t *Tree) Insert(val interface{}) {
	return t.i(t, val)
}

func containsInner(curNode Node, c Comparator, val interface{}) Node {
	if curNode == nil {
		return nil
	}
	v := c(val, curNode.Value())
	if v == 0 {
		return curNode
	}
	if v < 0 {
		return containsInner(curNode.Left(), c, val)
	}
	return containsInner(curNode.Right(), c, val)
}

func (t *Tree) Contains(val interface{}) Node {
	return containsInner(t.Root, t.C, val)
}

func (t *Tree) Remove(val interface{}) bool {
	return t.r(t, val)
}

func New(c Comparator, i InsertFunc, r RemoveFunc) *Tree {
	return &Tree{C: c, i: i, r:r, Root: nil}
}
