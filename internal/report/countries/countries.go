package countries

import "github.com/AdanJSuarez/maxmind/internal/node"

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

func (c *Countries) Countries() map[string]*node.Node {
	return c.countries.Children()
}
