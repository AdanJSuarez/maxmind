package countries

import (
	"github.com/AdanJSuarez/maxmind/internal/node"
)

const (
	minVisitors = 1
)

type Info struct {
	Name    string
	Visit   int64
	TopPage string
}

type countries struct {
	countries node.Node
}

// New returns a initialized instance of countries.
func New(name string) *countries {
	return &countries{
		countries: node.New(name),
	}
}

// Name returns countries name.
func (c *countries) Name() string {
	return c.countries.Name()
}

// AddToCountries adds a new element to countries.
func (c *countries) AddToCountries(countryName, subdivisionName, webpageName string) {
	c.countries.AddToNode(countryName, subdivisionName, webpageName)
}

// TopAreas returns the sorted Info about the area required by name.
// Area could be country, subdivision or others
func (c *countries) TopAreas(name, pageExcluded string, topNumber int) []Info {
	area := c.countries.FindNode(name)
	return c.topAreas(area, pageExcluded, topNumber)
}

// topAreas get the slice of children sorted by counter, then check if there is the
// minimum visits and then create a slice of info with the children Info up to a topNumber.
func (c *countries) topAreas(areas node.Node, pageExcluded string, topNumber int) []Info {
	var result []Info

	if areas == nil {
		return result
	}
	sortedChildren := areas.SortedChildren()
	for _, child := range sortedChildren {
		if c.notEnoughVisitors(child.Counter()) {
			continue
		}

		areaInfo := c.newInfo(child, pageExcluded, child.Counter())
		result = append(result, areaInfo)
		if len(result) >= topNumber {
			break
		}
	}
	return result
}

func (c *countries) notEnoughVisitors(visitors int64) bool {
	return visitors < minVisitors
}

func (c *countries) newInfo(area node.Node, pageExcluded string, visits int64) Info {
	return Info{
		Name:    area.Name(),
		Visit:   visits,
		TopPage: c.topPage(area, pageExcluded),
	}
}

func (c *countries) topPage(area node.Node, pageExcluded string) string {
	sortedData := area.SortedData(pageExcluded)
	if len(sortedData) == 0 {
		return ""
	}
	return sortedData[0].Name()
}
