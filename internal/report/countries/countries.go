package countries

import (
	"github.com/AdanJSuarez/maxmind/internal/node"
)

const (
	top          = 10
	oneVisitor   = 1
	unitedStates = "United States"
	countries    = "Countries"
	excluded     = "/"
)

type Info struct {
	Name    string
	Visit   int64
	TopPage string
}

type Countries struct {
	countries *node.Node
}

func New() *Countries {
	return &Countries{
		countries: node.New("Countries"),
	}
}

func (c *Countries) AddToCountries(countryName, subdivisionName, webpageName string) {
	c.countries.AddToNode(countryName, subdivisionName, webpageName)
}

func (c *Countries) Countries() *node.Node {
	return c.countries
}

func (c *Countries) TopAreas(name, excluded string) []Info {
	node := c.countries.FindNode(name)
	return c.topAreas(node, excluded)
}

func (c *Countries) topAreas(areas *node.Node, excluded string) []Info {
	var result []Info
	for _, val := range areas.SortedChildrenByCounter() {
		area, found := areas.Children()[val.Key()]
		if !found || val.Value() < oneVisitor {
			continue
		}
		areaInfo := Info{
			Name:    val.Key(),
			Visit:   val.Value(),
			TopPage: c.topPage(area, excluded),
		}
		result = append(result, areaInfo)
		if len(result) == 10 {
			break
		}
	}
	return result
}

func (c *Countries) topPage(area *node.Node, excluded string) string {
	if len(area.SortedData(excluded)) == 0 {
		return ""
	}
	return area.SortedData(excluded)[0].Key()
}
