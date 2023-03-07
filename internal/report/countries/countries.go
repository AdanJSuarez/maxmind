package countries

import (
	"github.com/AdanJSuarez/maxmind/internal/node"
)

const (
	oneVisitor = 1
	countries  = "Countries"
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
		countries: node.New(countries),
	}
}

func (c *Countries) AddToCountries(countryName, subdivisionName, webpageName string) {
	c.countries.AddToNode(countryName, subdivisionName, webpageName)
}

func (c *Countries) Countries() *node.Node {
	return c.countries
}

// TopAreas returns the sorted Info about the area required by name.
// Area could be country, subdivision or others
func (c *Countries) TopAreas(name, excluded string, topNumber int) []Info {
	area := c.countries.FindNode(name)
	return c.topAreas(area, excluded, topNumber)
}

func (c *Countries) topAreas(areas *node.Node, excluded string, topNumber int) []Info {
	var result []Info
	for _, child := range areas.SortedChildrenByCounter() {
		area, found := areas.Children()[child.Name()]
		if !found || c.notEnoughVisitors(child.Value()) {
			continue
		}

		areaInfo := c.newInfo(area, child.Name(), excluded, child.Value())
		result = append(result, areaInfo)
		if len(result) == topNumber {
			break
		}
	}
	return result
}

func (c *Countries) topPage(area *node.Node, excluded string) string {
	if len(area.SortedData(excluded)) == 0 {
		return ""
	}
	return area.SortedData(excluded)[0].Name()
}

func (c *Countries) notEnoughVisitors(value int64) bool {
	return value < oneVisitor
}

func (c *Countries) newInfo(area *node.Node, name, excluded string, visits int64) Info {
	return Info{
		Name:    name,
		Visit:   visits,
		TopPage: c.topPage(area, excluded),
	}
}
