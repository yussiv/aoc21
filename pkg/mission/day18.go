package mission

import (
	"fmt"

	"github.com/yussiv/aoc21/util"
)

type intNode struct {
	value  int
	depth  int
	left   *intNode
	right  *intNode
	parent *intNode
}

type Day18 struct {
	input []string
}

func (d *Day18) SetInput(input []string) {
	d.input = input
}

func (d *Day18) Task1() int {
	sum := d.buildSnailfishNumber(d.input[0])
	for _, numberString := range d.input[1:] {
		number := d.buildSnailfishNumber(numberString)
		sum = sum.add(number)
	}
	// fmt.Println(sum.toString())
	return sum.value
}

func (d *Day18) Task2() int {
	max := 0
	for i, numStr1 := range d.input {
		for j, numStr2 := range d.input {
			if i != j {
				num1 := d.buildSnailfishNumber(numStr1)
				num2 := d.buildSnailfishNumber(numStr2)
				num1 = num1.add(num2)
				if num1.value > max {
					max = num1.value
				}
			}
		}
	}
	return max
}

func (n *intNode) add(other *intNode) *intNode {
	newRoot := new(intNode)
	newRoot.left = n
	newRoot.right = other
	n.parent = newRoot
	other.parent = newRoot
	other.updateAncestors()
	newRoot.rebalance()
	return newRoot
}

func (n *intNode) rebalance() {
	n.explode()
	for toSplit := n.findSplittable(); toSplit != nil; toSplit = n.findSplittable() {
		toSplit.split()
		if n.depth > 4 {
			n.explode()
		}
	}
}

func (n *intNode) explode() {
	for n.depth > 4 {
		toExplode := n.findDeepest().parent
		nextLeft := toExplode.findNextLeft()
		nextRight := toExplode.findNextRight()
		if nextLeft != nil {
			nextLeft.value += toExplode.left.value
			nextLeft.updateAncestors()
		}
		if nextRight != nil {
			nextRight.value += toExplode.right.value
			nextRight.updateAncestors()
		}
		toExplode.value = 0
		toExplode.depth = 0
		toExplode.left = nil
		toExplode.right = nil
		toExplode.updateAncestors()
	}
}

func (n *intNode) findSplittable() *intNode {
	if n.left == nil {
		if n.value > 9 {
			return n
		}
		return nil
	}
	left := n.left.findSplittable()
	if left == nil {
		return n.right.findSplittable()
	}
	return left
}

func (n *intNode) split() {
	left := new(intNode)
	right := new(intNode)
	left.parent = n
	right.parent = n
	left.value = n.value / 2
	right.value = n.value - n.value/2
	n.left = left
	n.right = right
	left.updateAncestors()
}

func (n *intNode) updateAncestors() {
	parent := n.parent
	for parent != nil {
		parent.depth = util.Max(parent.left.depth, parent.right.depth) + 1
		parent.value = 3*parent.left.value + 2*parent.right.value
		parent = parent.parent
	}
}

func (n *intNode) findDeepest() *intNode {
	node := n
	for node.depth != 0 {
		if node.left.depth >= node.right.depth {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

func (n *intNode) findNextLeft() *intNode {
	if n.parent == nil {
		return nil
	}
	node := n
	for {
		if node.parent.left != node {
			node = node.parent.left
			break
		} else if node.parent.parent == nil {
			return nil
		} else {
			node = node.parent
		}
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

func (n *intNode) findNextRight() *intNode {
	if n.parent == nil {
		return nil
	}
	node := n
	for {
		if node.parent.right != node {
			node = node.parent.right
			break
		} else if node.parent.parent == nil {
			return nil
		} else {
			node = node.parent
		}
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func (n intNode) toString() string {
	if n.left == nil || n.right == nil {
		return fmt.Sprint(n.value)
	}
	return fmt.Sprintf("[%s,%s]", n.left.toString(), n.right.toString())
}

func (d Day18) buildSnailfishNumber(input string) *intNode {
	root, _ := d.buildNode([]rune(input), 0, nil)
	return root
}

func (d Day18) buildNode(input []rune, i int, parent *intNode) (*intNode, int) {
	node := new(intNode)
	node.parent = parent

	for i < len(input) {
		switch r := input[i]; r {
		case '[':
			node.left, i = d.buildNode(input, i+1, node)
		case ',':
			node.right, i = d.buildNode(input, i+1, node)
		case ']':
			node.depth = util.Max(node.left.depth, node.right.depth) + 1
			node.value = 3*node.left.value + 2*node.right.value
			return node, i + 1
		default:
			node.value = int(r - '0')
			return node, i + 1
		}
	}
	fmt.Println("whoops, that's not right")
	return node, i + 1
}
