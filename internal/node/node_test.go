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
