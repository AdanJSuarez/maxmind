package node

import "sort"

type node struct {
	name     string
	counter  int64
	data     map[string]int64
	children map[string]Node
}

// New returns a initialized instance of an object that implement Node.
func New(name string) Node {
	return &node{
		name:     name,
		data:     make(map[string]int64),
		children: make(map[string]Node),
	}
}

// Getters

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

// SortedData returns a slice of the data sorted by its counter.
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

// SortedChildrenByCounter returns a slice of Children sorted by its counter.
func (n *node) SortedChildren() []Node {
	sorted := make([]Node, 0, len(n.children))

	for _, val := range n.children {
		sorted = append(sorted, val)
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Counter() > sorted[j].Counter()
	})

	return sorted
}

// FindNode makes a Breadth-First Search and returns the first Node found.
func (n *node) FindNode(name string) Node {
	todo := []Node{}
	visited := []Node{}
	return n.bfsForNode(name, n, todo, visited)
}

// addToData adds data to node.data and increases its counter
func (n *node) addToData(data string) {
	n.data[data]++
}

// addToChild adds element to children.
func (n *node) addToChild(parameters ...string) {
	childName := parameters[0]
	_, found := n.children[childName]
	if !found {
		node := New(childName)
		n.children[childName] = node
	}
	n.children[childName].AddToNode(removeFirstElement(parameters)...)
}

func (n *node) hasOneOrMoreElement(parametersLen int) bool {
	return parametersLen > 0
}

func (n *node) hasTwoOrMoreElement(parametersLen int) bool {
	return parametersLen > 1
}

// bfsForNode is the recursive function for node to traverse the tree (BFS)
func (n *node) bfsForNode(name string, node Node, todo []Node, visited []Node) Node {
	if n.contain(node, visited) {
		return nil
	}

	if node.Name() == name {
		return node
	}

	for _, val := range node.Children() {
		todo = append(todo, val)
	}

	return n.bfsForTodo(name, todo, visited)
}

// bfsForTodo is the recursive function for Todo to traverse the tree (BFS)
func (n *node) bfsForTodo(name string, todo []Node, visited []Node) Node {
	if len(todo) == 0 {
		return nil
	}

	return n.bfsForNode(name, todo[0], removeFirstElement(todo), visited)
}

func (n *node) contain(node Node, nodes []Node) bool {
	for _, n := range nodes {
		if node.Name() == n.Name() {
			return true
		}
	}
	return false
}

// removeFirstElement returns the slice with the first element removed.
func removeFirstElement[T any](values []T) []T {
	if len(values) == 0 {
		return values
	}
	return values[1:]
}
