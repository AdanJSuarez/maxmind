package node

import "sort"

type Node struct {
	name     string
	counter  int64
	data     map[string]int64
	children map[string]*Node
}

// TODO: Move Data to its own file
type Data struct {
	key   string
	value int64
}

func (d *Data) Key() string {
	return d.key
}

func (d *Data) Value() int64 {
	return d.value
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

// SortedData returns a sorted slice of the data.
func (n *Node) SortedData(excluded string) []Data {
	sorted := make([]Data, 0, len(n.data))
	for key, val := range n.data {
		if key == excluded {
			continue
		}
		sorted = append(sorted, Data{key, val})
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].value > sorted[j].value
	})

	return sorted
}

// SortedChildrenByCounter returns a sorted slice of Children by its counter
func (n *Node) SortedChildrenByCounter() []Data {
	sorted := make([]Data, 0, len(n.data))
	for key, val := range n.children {
		sorted = append(sorted, Data{key, val.counter})
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].value > sorted[j].value
	})

	return sorted
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

// Function for node:
// Read a node with todo, visited.
// If node visited return
// Else:
// Add to visited.
// Add children to todos.
// Compute on that node.
// Call function for todos.

// Function for todo:
// If todo is empty, return
// If todo not empty:
// Function for node[0] with removed from todo.
