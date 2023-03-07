package node

//go:generate mockery --inpackage --name=Node

type Node interface {
	Name() string
	Counter() int64
	Data() map[string]int64
	Children() map[string]Node
	SortedChildren() []Node
	SortedData(pageExcluded string) []Data
	AddToNode(parameters ...string)
	FindNode(name string) Node
}
