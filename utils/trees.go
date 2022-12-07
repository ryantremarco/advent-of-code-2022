package utils

type Tree[N comparable, L any] struct {
	*Node[N, L]
}

func NewTree[N comparable, L any](headValue N) Tree[N, L] {
	return Tree[N, L]{
		Node: NewNode[N, L](headValue, nil),
	}
}

type Node[N comparable, L any] struct {
	Value      N
	ChildNodes []*Node[N, L]
	Leaves     []L
	Parent     *Node[N, L]
}

func NewNode[N comparable, L any](value N, parent *Node[N, L]) *Node[N, L] {
	return &Node[N, L]{
		Value:      value,
		ChildNodes: []*Node[N, L]{},
		Leaves:     []L{},
		Parent:     parent,
	}
}

func (n *Node[N, L]) AddNode(toAdd N) *Node[N, L] {
	newNode := NewNode(toAdd, n)
	n.ChildNodes = append(n.ChildNodes, newNode)
	return newNode
}

func (n *Node[N, L]) GetNode(value N) *Node[N, L] {
	for _, node := range n.ChildNodes {
		if node != nil && node.Value == value {
			return node
		}
	}
	return nil
}

func (n *Node[N, L]) AllLeaves() []L {
	leaves := append([]L{}, n.Leaves...)
	for _, node := range n.ChildNodes {
		leaves = append(leaves, node.AllLeaves()...)
	}
	return leaves
}

func (n *Node[N, L]) AddLeaf(toAdd L) {
	n.Leaves = append(n.Leaves, toAdd)
}
