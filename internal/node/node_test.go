package node

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	nodeTest Node
)

type TSNode struct{ suite.Suite }

func TestRunTSNode(t *testing.T) {
	suite.Run(t, new(TSNode))
}

func (ts *TSNode) BeforeTest(_, _ string) {
	nodeTest = New("Spain")
}

func (ts *TSNode) TestNodeInitialized() {
	ts.Equal("Spain", nodeTest.Name())
	ts.NotNil(nodeTest.Data())
	ts.NotNil(nodeTest.Children())
}

func (ts *TSNode) TestAddToNode() {
	nodeTest.AddToNode("Tenerife", "/")
	pageCounter := nodeTest.Data()["/"]
	ts.Equal(int64(1), pageCounter)

	child := nodeTest.Children()["Tenerife"]
	ts.Equal("Tenerife", child.Name())
	ts.Equal(int64(1), child.Counter())

	grantChild := nodeTest.Children()["Tenerife"].Children()
	ts.Empty(grantChild)
}

func (ts *TSNode) TestAddToNodeTwo() {
	nodeTest.AddToNode("Tenerife", "/")
	nodeTest.AddToNode("Tenerife", "/")

	pageCounter := nodeTest.Data()["/"]
	ts.Equal(int64(2), pageCounter)

	child := nodeTest.Children()["Tenerife"]
	ts.Equal("Tenerife", child.Name())
	ts.Equal(int64(2), child.Counter())

	grantChild := nodeTest.Children()["Tenerife"].Children()
	ts.Empty(grantChild)
}

func (ts *TSNode) TestAddToNodeTwoDifferent() {
	nodeTest.AddToNode("Tenerife", "/")
	nodeTest.AddToNode("Tenerife", "/turbo")

	pageCounter1 := nodeTest.Data()["/"]
	ts.Equal(int64(1), pageCounter1)

	pageCounter2 := nodeTest.Data()["/turbo"]
	ts.Equal(int64(1), pageCounter2)

	child := nodeTest.Children()["Tenerife"]
	ts.Equal("Tenerife", child.Name())
	ts.Equal(int64(2), child.Counter())

	ts.Equal(int64(1), child.Data()["/"])
	ts.Equal(int64(1), child.Data()["/turbo"])

	grantChild := nodeTest.Children()["Tenerife"].Children()
	ts.Empty(grantChild)
}

func (ts *TSNode) TestAddToNodeTwoChild() {
	nodeTest.AddToNode("Madrid", "/")
	nodeTest.AddToNode("Tenerife", "/turbo")

	pageCounter1 := nodeTest.Data()["/"]
	ts.Equal(int64(1), pageCounter1)

	pageCounter2 := nodeTest.Data()["/turbo"]
	ts.Equal(int64(1), pageCounter2)

	child1 := nodeTest.Children()["Tenerife"]
	ts.Equal("Tenerife", child1.Name())
	ts.Equal(int64(1), child1.Counter())

	ts.Equal(int64(0), child1.Data()["/"])
	ts.Equal(int64(1), child1.Data()["/turbo"])

	child2 := nodeTest.Children()["Madrid"]
	ts.Equal("Madrid", child2.Name())
	ts.Equal(int64(1), child2.Counter())

	ts.Equal(int64(1), child2.Data()["/"])
	ts.Equal(int64(0), child2.Data()["/turbo"])

	grantChild := nodeTest.Children()["Tenerife"].Children()
	ts.Empty(grantChild)
}

func (ts *TSNode) TestAddToNodeGrandChild() {
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")

	pageCounter := nodeTest.Data()["/"]
	ts.Equal(int64(1), pageCounter)

	child := nodeTest.Children()["Tenerife"]
	ts.Equal("Tenerife", child.Name())
	ts.Equal(int64(1), child.Counter())

	grantChild := nodeTest.Children()["Tenerife"].Children()
	ts.Equal("Santa Ursula", grantChild["Santa Ursula"].Name())
	ts.Equal(int64(1), grantChild["Santa Ursula"].Counter())
}

func (ts *TSNode) TestAddToNodeGrandChild2() {
	nodeTest.AddToNode("Tenerife", "La Matanza", "/turbo")
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")
	nodeTest.AddToNode("Tenerife", "La Victoria", "/")

	nodeTest.AddToNode("Canada", "Ontario", "/turbo1")
	nodeTest.AddToNode("Canada", "Alberta", "/t")
	nodeTest.AddToNode("Canada", "Alberta", "/t")
	nodeTest.AddToNode("Canada", "Alberta", "/trading")
	nodeTest.AddToNode("Canada", "Alberta", "/t")
	nodeTest.AddToNode("Canada", "Yukon", "/t")

	pageCounter := nodeTest.Data()["/"]
	ts.Equal(int64(2), pageCounter)

	canada := nodeTest.Children()["Canada"]
	ts.Equal(int64(4), canada.Data()["/t"])

	alberta := canada.Children()["Alberta"]
	ts.Equal(int64(4), alberta.Counter())
	ts.Equal(int64(3), alberta.Data()["/t"])
	ts.Equal(int64(0), alberta.Data()["/turbo1"])
}

func (ts *TSNode) TestSortedData() {
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")
	nodeTest.AddToNode("Tenerife", "La Victoria", "/")
	nodeTest.AddToNode("Tenerife", "La Matanza", "/")

	actual := nodeTest.SortedData("")
	expected := []Data{{name: "/", counter: 3}}
	ts.Equal(expected, actual)
}
func (ts *TSNode) TestSortedData2() {
	nodeTest.AddToNode("Tenerife", "La Matanza", "/turbo")
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")
	nodeTest.AddToNode("Tenerife", "La Victoria", "/")

	actual := nodeTest.SortedData("")
	expected := []Data{
		{name: "/", counter: 2},
		{name: "/turbo", counter: 1},
	}
	ts.Equal(expected, actual)
}
func (ts *TSNode) TestSortedDataPageExcluded() {
	nodeTest.AddToNode("Tenerife", "La Matanza", "/turbo")
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")
	nodeTest.AddToNode("Tenerife", "La Victoria", "/")

	actual := nodeTest.SortedData("/")
	expected := []Data{
		{name: "/turbo", counter: 1},
	}
	ts.Equal(expected, actual)
}
func (ts *TSNode) TestSortedDataEmptyData() {
	actual := nodeTest.SortedData("/")
	expected := []Data{}
	ts.Equal(expected, actual)
}

func (ts *TSNode) TestSortedChildrenByCounter() {
	node1 := &node{name: "La Matanza", counter: 1}
	node2 := &node{name: "Santa Ursula", counter: 2}
	node3 := &node{name: "La Victoria", counter: 3}

	nodeTest := node{name: "countries", children: map[string]Node{}}
	nodeTest.children[node1.name] = node1
	nodeTest.children[node2.name] = node2
	nodeTest.children[node3.name] = node3

	actual := nodeTest.SortedChildren()
	expected := []Node{node3, node2, node1}
	ts.Equal(expected, actual)
}

func (ts *TSNode) TestSortedChildrenByCounterEmpty() {
	nodeTest := node{name: "countries", children: map[string]Node{}}

	actual := nodeTest.SortedChildren()
	expected := []Node{}
	ts.Equal(expected, actual)
}

func (ts *TSNode) TestRemoveFirstElement() {
	names := []string{"a"}
	actual := removeFirstElement(names)
	ts.Empty(actual)
}
func (ts *TSNode) TestRemoveFirstElement2() {
	names := []string{}
	actual := removeFirstElement(names)
	ts.Empty(actual)
}

func (ts *TSNode) TestRemoveFirstElement3() {
	names := []string{"a", "b", "c"}
	actual := removeFirstElement(names)
	ts.Equal([]string{"b", "c"}, actual)
}

func (ts *TSNode) TestFindNode1() {
	nodeTest.AddToNode("Tenerife", "La Matanza", "/turbo")
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")
	nodeTest.AddToNode("Tenerife", "La Victoria", "/")

	nodeTest.AddToNode("Canada", "Ontario", "/turbo1")
	nodeTest.AddToNode("Canada", "Alberta", "/t")
	nodeTest.AddToNode("Canada", "Yukon", "/x")

	actual := nodeTest.FindNode("Canada")
	ts.NotNil(actual)
	child, found := actual.Children()["Alberta"]
	ts.Equal(child.Name(), "Alberta")
	ts.True(found)
	ts.Equal(int64(1), child.Counter())
}

func (ts *TSNode) TestFindNode2() {
	nodeTest.AddToNode("Tenerife", "La Matanza", "/turbo")
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")
	nodeTest.AddToNode("Tenerife", "La Victoria", "/")

	nodeTest.AddToNode("Canada", "Ontario", "/turbo1")
	nodeTest.AddToNode("Canada", "Alberta", "/t")
	nodeTest.AddToNode("Canada", "Yukon", "/x")

	actual := nodeTest.FindNode("Canada")
	ts.NotNil(actual)
	child, found := actual.Children()["Manitoba"]
	ts.Nil(child)
	ts.False(found)
}

func (ts *TSNode) TestFindNode3() {
	nodeTest.AddToNode("Tenerife", "La Matanza", "/turbo")
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")
	nodeTest.AddToNode("Tenerife", "La Victoria", "/")

	nodeTest.AddToNode("Canada", "Ontario", "/turbo1")
	nodeTest.AddToNode("Canada", "Alberta", "/t")
	nodeTest.AddToNode("Canada", "Yukon", "/x")

	actual := nodeTest.FindNode("Alberta")
	ts.NotNil(actual)
	child, found := actual.Children()["Calgary"]
	ts.Nil(child)
	ts.False(found)
}

func (ts *TSNode) TestFindNode4() {
	nodeTest.AddToNode("Tenerife", "La Matanza", "/turbo")
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")
	nodeTest.AddToNode("Tenerife", "La Victoria", "/")

	nodeTest.AddToNode("Canada", "Ontario", "/turbo1")
	nodeTest.AddToNode("Canada", "Alberta", "/t")
	nodeTest.AddToNode("Canada", "Yukon", "/x")

	actual := nodeTest.FindNode("Montana")
	ts.Nil(actual)
}

func (ts *TSNode) TestNodeVisitedAlready() {
	node := &node{name: "La Matanza", counter: 1}
	visited := []Node{node}
	todo := []Node{}

	actual := node.bfsForNode("La Matanza", node, todo, visited)
	ts.Nil(actual)
}

func (ts *TSNode) TestContain() {
	node1 := &node{name: "La Matanza", counter: 1}
	node2 := &node{name: "Santa Ursula", counter: 2}
	node3 := &node{name: "La Victoria", counter: 3}

	nodes := []Node{node1, node2, node3}

	actual := node1.contain(node3, nodes)
	ts.True(actual)
}
func (ts *TSNode) TestContain2() {
	node1 := &node{name: "La Matanza", counter: 1}
	node2 := &node{name: "Santa Ursula", counter: 2}
	node3 := &node{name: "La Victoria", counter: 3}

	nodes := []Node{node1, node2}

	actual := node1.contain(node3, nodes)
	ts.False(actual)
}
