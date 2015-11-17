package node

type Node struct {
	x        int
	y        int
	step     int
	children []*Node
}

func (node *Node) GetX() int {
	return node.x
}
func (node *Node) SetX(x int) {
	node.x = x
}

func (node *Node) GetY() int {
	return node.y
}
func (node *Node) SetY(y int) {
	node.y = y
}

func (node *Node) GetStep() int {
	return node.step
}
func (node *Node) SetStep(step int) {
	node.step = step
}

func (node *Node) AddChild(child *Node) {
	node.children = append(node.children, child)
}
