package countries

import (
	"testing"

	"github.com/AdanJSuarez/maxmind/internal/node"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const countriesNameTest = "Countries"

var (
	countriesTest *countries
	nodeMock      *node.MockNode
	mockChild     *node.MockNode
)

type TSCountries struct{ suite.Suite }

func TestRunTSCountries(t *testing.T) {
	suite.Run(t, new(TSCountries))
}

func (ts *TSCountries) BeforeTest(_, _ string) {
	nodeMock = node.NewMockNode(ts.T())
	mockChild = node.NewMockNode(ts.T())
	countriesTest = New(countriesNameTest)
	countriesTest.countries = nodeMock
}

func (ts *TSCountries) TestCountriesInitialization() {
	ts.NotNil(countriesTest.countries)
}

func (ts *TSCountries) TestName() {
	nodeMock.On("Name", mock.Anything).Return(countriesNameTest)
	name := countriesTest.Name()
	ts.Equal("Countries", name)
}

func (ts *TSCountries) TestAddToCountries1() {
	nodeMock.On("AddToNode", mock.Anything, mock.Anything, mock.Anything).Return()
	countriesTest.AddToCountries("United States", "Montana", "/")

	ts.True(nodeMock.AssertNumberOfCalls(ts.T(), "AddToNode", 1))
}

func (ts *TSCountries) TestTopAreasAtLestOneVisitor() {
	mockChild.On("Counter").Return(int64(1))
	mockChild.On("Name").Return("Tenerife")
	mockChild.On("SortedData", mock.Anything).Return([]node.Data{node.NewData("/", 1)})
	nodeMock.On("FindNode", mock.Anything).Return(nodeMock)
	nodeMock.On("SortedChildren").Return([]node.Node{mockChild})

	info := countriesTest.TopAreas("Countries,", "", 10)
	ts.Equal([]Info{{Name: "Tenerife", Visit: 1, TopPage: "/"}}, info)
}

func (ts *TSCountries) TestTopAreasAtTwoVisitors() {
	mockChild.On("Counter").Return(int64(2))
	mockChild.On("Name").Return("Tenerife")
	mockChild.On("SortedData", mock.Anything).Return([]node.Data{node.NewData("/", 2)})
	nodeMock.On("FindNode", mock.Anything).Return(nodeMock)
	nodeMock.On("SortedChildren").Return([]node.Node{mockChild})

	info := countriesTest.TopAreas("Countries,", "", 10)
	expected := []Info{
		{Name: "Tenerife", Visit: 2, TopPage: "/"},
	}
	ts.Equal(expected, info)
}

func (ts *TSCountries) TestTopAreasNoVisitor() {
	mockChild.On("Counter").Return(int64(0))
	nodeMock.On("FindNode", mock.Anything).Return(nodeMock)
	nodeMock.On("SortedChildren").Return([]node.Node{mockChild})

	info := countriesTest.TopAreas("Countries,", "", 10)
	ts.Empty(info)
}

func (ts *TSCountries) TestTopAreasSelectNoMoreThanTopNumber() {
	mockChild.On("Counter").Return(int64(2))
	mockChild.On("Name").Return("Tenerife")
	mockChild.On("SortedData", mock.Anything).Return([]node.Data{node.NewData("/", 2)})
	nodeMock.On("FindNode", mock.Anything).Return(nodeMock)
	nodeMock.On("SortedChildren").Return([]node.Node{mockChild, mockChild, mockChild})

	info := countriesTest.TopAreas("Countries,", "", 2)
	expected := []Info{
		{Name: "Tenerife", Visit: 2, TopPage: "/"},
		{Name: "Tenerife", Visit: 2, TopPage: "/"},
	}
	ts.Equal(expected, info)
}

func (ts *TSCountries) TestTopPage() {
	data1 := node.NewData("/", 3)
	data2 := node.NewData("/turbo", 2)
	data := []node.Data{data1, data2}
	nodeMock.On("SortedData", mock.Anything).Return(data)

	page := countriesTest.topPage(nodeMock, "")
	ts.Equal("/", page)
}

func (ts *TSCountries) TestNoTopPage() {
	data := []node.Data{}
	nodeMock.On("SortedData", mock.Anything).Return(data)

	page := countriesTest.topPage(nodeMock, "")
	ts.Equal("", page)
}
