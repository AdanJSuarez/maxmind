package node

import "sort"

type node struct {
	name     string
	counter  int64
	data     map[string]int64
	children map[string]Node
}

// New returns a initialized instance of Node
func New(name string) Node {
	return &node{
		name:     name,
		data:     make(map[string]int64),
		children: make(map[string]Node),
	}
}

func (n *node) Name() string {
	return n.name
}

func (n *node) Counter() int64 {
	return n.counter
}

func (n *node) Data() map[string]int64 {
	return n.data
}

func (n *node) Children() map[string]Node {
	return n.children
}

// AddToNode adds data to the node and their children if applicable.
// It increases the counter with each adding.
func (n *node) AddToNode(parameters ...string) {
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
func (n *node) SortedData(pageExcluded string) []Data {
	sorted := make([]Data, 0, len(n.data))
	for key, counter := range n.data {
		if key == pageExcluded {
			continue
		}
		sorted = append(sorted, NewData(key, counter))
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].counter > sorted[j].counter
	})

	return sorted
}

// SortedChildrenByCounter returns a sorted slice of Children by its counter
func (n *node) SortedChildrenByCounter() []Node {
	sorted := make([]Node, 0, len(n.data))
	for _, val := range n.children {
		sorted = append(sorted, val)
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Counter() > sorted[j].Counter()
	})

	return sorted
}

// FindNode makes a Breadth-First Search and return the first node with "name"
func (n *node) FindNode(name string) Node {
	todo := make(map[string]Node)
	visited := make(map[string]Node)
	return n.traverseForNode(name, n, todo, visited)
}

func (n *node) traverseForNode(name string, node Node, todo map[string]Node, visited map[string]Node) Node {
	if n.contain(node, visited) {
		return nil
	}
	if node.Name() == name {
		return node
	}
	for key, val := range node.Children() {
		todo[key] = val
	}
	return n.traverseForTodo(name, todo, visited)
}

func (n *node) traverseForTodo(name string, todo map[string]Node, visited map[string]Node) Node {
	if len(todo) == 0 {
		return nil
	}
	var node Node
	var nodeName string
	for key, val := range todo {
		node = val
		nodeName = key
		break
	}
	delete(todo, nodeName)
	return n.traverseForNode(name, node, todo, visited)
}

// addToData adds data to node.data and increases its counter
func (n *node) addToData(data string) {
	n.data[data]++
}

func (n *node) addToChild(parameters ...string) {
	childName := parameters[0]
	_, found := n.children[childName]
	if !found {
		node := New(childName)
		n.children[childName] = node
	}
	n.children[childName].AddToNode(n.removeFirstElement(parameters)...)
}

func (n *node) hasOneOrMoreElement(parametersLen int) bool {
	return parametersLen > 0
}

func (n *node) hasTwoOrMoreElement(parametersLen int) bool {
	return parametersLen > 1
}

func (n *node) removeFirstElement(parameters []string) []string {
	return parameters[1:]
}

func (n *node) contain(node Node, nodes map[string]Node) bool {
	_, found := nodes[node.Name()]
	return found
}
