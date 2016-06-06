package binary

import (
	"fmt"
	"github.com/jonbodner/trees"
	"math"
	"math/rand"
)

type nodeBinary struct {
	value               interface{}
	left, right, parent *nodeBinary
}

func (n *nodeBinary) Value() interface{} {
	return n.value
}

func (n *nodeBinary) Left() trees.Node {
	return n.left
}

func (n *nodeBinary) Right() trees.Node {
	return n.right
}

func (n *nodeBinary) setLeft(in nodeBinary) {
	n.left = in
}

func (n *nodeBinary) setRight(in nodeBinary) {
	n.right = in
}

func printInner(n trees.Node, level int, pos int, vals *[][]interface{}) {
	if len(*vals) == level {
		curRow := make([]interface{}, int(math.Pow(2.0, float64(level))))
		*vals = append(*vals, curRow)
	}
	curRow := (*vals)[level]
	curRow[pos] = n.Value()
	if n.Left() != nil {
		printInner(n.Left(), level+1, pos*2, vals)
	}
	if n.Right() != nil {
		printInner(n.Right(), level+1, pos*2+1, vals)
	}
}

func (n *nodeBinary) String() string {
	vals := [][]interface{}{}
	printInner(n, 0, 0, &vals)
	s := ""
	for _, v := range vals {
		s += fmt.Sprintf("%v\n", v)
	}
	return s
}

func Insert(t *trees.Tree, val interface{}) {
	x := &nodeBinary{value:val}
	if t.Root == nil {
		t.Root = x
		return
	}
	curNode := t.Root.(*nodeBinary)
	for {
		v := t.C(x.Value(), curNode.Value())
		if v == 0 {
			//random choice -- make it -1 or 1 and continue
			if rand.Intn(2) == 0 {
				v = -1
			} else {
				v = 1
			}
		}
		if v < 0 {
			if curNode.Left() == nil {
				curNode.setLeft(x)
				return
			}
			curNode = curNode.Left().(*nodeBinary)
		} else {
			if curNode.Right() == nil {
				curNode.setRight(x)
				return
			}
			curNode = curNode.Right().(*nodeBinary)
		}
	}
}

func Remove(t *trees.Tree, val interface{}) bool {
	//find the node
	curNode := t.Contains(val)
	//if not present, return false
	if curNode == nil {
		return false
	}
	curNodeBinary := curNode.(*nodeBinary)
	//if only a left node or no children, make it the right or left of my parent (whichever I was) and return true
	if curNodeBinary.Right() == nil {
		if t.Root == curNodeBinary {
			t.Root = curNodeBinary.left
		} else if curNodeBinary.parent.left == curNodeBinary {
			curNodeBinary.parent.left = curNodeBinary.left
		} else {
			curNodeBinary.parent.right = curNodeBinary.left
		}
		if curNodeBinary.left != nil {
			curNodeBinary.left.parent = curNodeBinary.parent
		}
		return true
	}
	//take my right node
	//make it the right or left of my parent (whichever I was)
	//take my left node
	//if non null,
		// find the leftmost descendant of my right node that has no left child
		//make my left node the left node of that descendant
	//return true
}