package countries

import "github.com/AdanJSuarez/maxmind/internal/node"

type Countries struct {
	countries map[string]*node.Node
}

func New() *Countries {
	return &Countries{
		countries: make(map[string]*node.Node),
	}
}

func (c *Countries) AddToCountries(countryName, subdivisionName, webpageName string) {
	_, found := c.countries[countryName]
	if !found {
		country := node.New(countryName)
		c.countries[countryName] = country
	}

	c.countries[countryName].AddToNode(subdivisionName, webpageName)
}

func (c *Countries) Countries() map[string]*node.Node {
	return c.countries
}
