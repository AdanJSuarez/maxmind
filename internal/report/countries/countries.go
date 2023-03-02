package countries

type Countries struct {
	countries map[string]*Country
}

func NewCountries() *Countries {
	return &Countries{
		countries: make(map[string]*Country),
	}
}

func (c *Countries) AddToCountries(countryName, subdivisionName, webpageName string) {
	_, found := c.countries[countryName]
	if !found {
		c.countries[countryName] = newCountry(countryName, subdivisionName, webpageName)
		return
	}
	c.countries[countryName].addToCountry(subdivisionName, webpageName)
}

func (c *Countries) Countries() map[string]*Country {
	return c.countries
}

func (c *Countries) MostVisit(number int) []*Country {
	result := make([]*Country, number)
	for _, v := range c.countries {
		for idx := range result {
			if result[idx] == nil {
				result[idx] = v
				break
			}
			if result[idx].counter <= v.counter {
				result[idx] = v
				break
			}
		}
	}
	return result
}
