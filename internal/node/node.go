package node

type Node struct {
	name     string
	counter  int64
	data     map[string]int64
	children map[string]*Node
}

// New returns a initialized instance of Node
func New(name string) *Node {
	return &Node{
		name:     name,
		data:     make(map[string]int64),
		children: make(map[string]*Node),
	}
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Counter() int64 {
	return n.counter
}

func (n *Node) Data() map[string]int64 {
	return n.data
}

func (n *Node) Children() map[string]*Node {
	return n.children
}

// AddToNode adds data to the node and their children if applicable.
// It increases the counter with each adding.
func (n *Node) AddToNode(parameters ...string) {
	parametersLen := len(parameters)
	var data string

	if n.hasOneOrMoreElement(parametersLen) {
		data = parameters[parametersLen-1]
	}

	n.counter++
	n.addToData(data)

	if n.hasTwoOrMoreElement(parametersLen) {
		n.addToChild(parameters...)
	}
}

func (n *Node) MostData(excluded string) {
	var max int64
	for data, counter := range n.data {
		if data != excluded && counter > max {
			max = counter
		}
	}
}

// addToData adds data to node.data and increases its counter
func (n *Node) addToData(data string) {
	n.data[data]++
}

func (n *Node) addToChild(parameters ...string) {
	childName := parameters[0]
	_, found := n.children[childName]
	if !found {
		node := New(childName)
		n.children[childName] = node
	}
	n.children[childName].AddToNode(n.removeFirstElement(parameters)...)
}

func (n *Node) hasOneOrMoreElement(parametersLen int) bool {
	return parametersLen > 0
}

func (n *Node) hasTwoOrMoreElement(parametersLen int) bool {
	return parametersLen > 1
}

func (n *Node) removeFirstElement(parameters []string) []string {
	return parameters[1:]
}
