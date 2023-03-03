package node

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	nodeTest *Node
)

type TSNode struct{ suite.Suite }

func TestRunTSNode(t *testing.T) {
	suite.Run(t, new(TSNode))
}

func (ts *TSNode) BeforeTest(_, _ string) {
	nodeTest = New("Spain")
}

func (ts *TSNode) TestNodeInitialized() {
	ts.Equal("Spain", nodeTest.name)
	ts.NotNil(nodeTest.data)
	ts.NotNil(nodeTest.children)
}

func (ts *TSNode) TestAddToNode() {
	nodeTest.AddToNode("Tenerife", "/")
	pageCounter := nodeTest.data["/"]
	ts.Equal(int64(1), pageCounter)

	child := nodeTest.children["Tenerife"]
	ts.Equal("Tenerife", child.name)
	ts.Equal(int64(1), child.counter)

	grantChild := nodeTest.children["Tenerife"].children
	ts.Empty(grantChild)
}

func (ts *TSNode) TestAddToNodeTwo() {
	nodeTest.AddToNode("Tenerife", "/")
	nodeTest.AddToNode("Tenerife", "/")

	pageCounter := nodeTest.data["/"]
	ts.Equal(int64(2), pageCounter)

	child := nodeTest.children["Tenerife"]
	ts.Equal("Tenerife", child.name)
	ts.Equal(int64(2), child.counter)

	grantChild := nodeTest.children["Tenerife"].children
	ts.Empty(grantChild)
}

func (ts *TSNode) TestAddToNodeTwoDifferent() {
	nodeTest.AddToNode("Tenerife", "/")
	nodeTest.AddToNode("Tenerife", "/turbo")

	pageCounter1 := nodeTest.data["/"]
	ts.Equal(int64(1), pageCounter1)

	pageCounter2 := nodeTest.data["/turbo"]
	ts.Equal(int64(1), pageCounter2)

	child := nodeTest.children["Tenerife"]
	ts.Equal("Tenerife", child.name)
	ts.Equal(int64(2), child.counter)

	ts.Equal(int64(1), child.data["/"])
	ts.Equal(int64(1), child.data["/turbo"])

	grantChild := nodeTest.children["Tenerife"].children
	ts.Empty(grantChild)
}

func (ts *TSNode) TestAddToNodeTwoChild() {
	nodeTest.AddToNode("Madrid", "/")
	nodeTest.AddToNode("Tenerife", "/turbo")

	pageCounter1 := nodeTest.data["/"]
	ts.Equal(int64(1), pageCounter1)

	pageCounter2 := nodeTest.data["/turbo"]
	ts.Equal(int64(1), pageCounter2)

	child1 := nodeTest.children["Tenerife"]
	ts.Equal("Tenerife", child1.name)
	ts.Equal(int64(1), child1.counter)

	ts.Equal(int64(0), child1.data["/"])
	ts.Equal(int64(1), child1.data["/turbo"])

	child2 := nodeTest.children["Madrid"]
	ts.Equal("Madrid", child2.name)
	ts.Equal(int64(1), child2.counter)

	ts.Equal(int64(1), child2.data["/"])
	ts.Equal(int64(0), child2.data["/turbo"])

	grantChild := nodeTest.children["Tenerife"].children
	ts.Empty(grantChild)
}

func (ts *TSNode) TestAddToNodeGrandChild() {
	nodeTest.AddToNode("Tenerife", "Santa Ursula", "/")

	pageCounter := nodeTest.data["/"]
	ts.Equal(int64(1), pageCounter)

	child := nodeTest.children["Tenerife"]
	ts.Equal("Tenerife", child.name)
	ts.Equal(int64(1), child.counter)

	grantChild := nodeTest.children["Tenerife"].children
	ts.Equal("Santa Ursula", grantChild["Santa Ursula"].name)
	ts.Equal(int64(1), grantChild["Santa Ursula"].counter)
}
