package countries

import (
	"github.com/AdanJSuarez/maxmind/internal/node"
)

const (
	top          = 10
	unitedStates = "United States"
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

func (c *Countries) TopPages(excluded string) []node.Data {
	return c.countries.SortedData(excluded)[:top]
}

func (c *Countries) TopAreas(areas *node.Node, exclude string) []Info {
	var result []Info
	for _, val := range areas.SortedChildrenByCounter() {
		area, found := areas.Children()[val.Key()]
		if !found {
			continue
		}
		areaInfo := Info{
			Name:    val.Key(),
			Visit:   val.Value(),
			TopPage: area.SortedData(exclude)[0].Key(),
		}
		result = append(result, areaInfo)
		if len(result) == 10 {
			break
		}
	}
	return result
}
