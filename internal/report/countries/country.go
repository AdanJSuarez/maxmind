package countries

type Country struct {
	name         string
	counter      int64
	subdivisions map[string]*Subdivision
}

func (c *Country) Name() string {
	return c.name
}

func (c *Country) Counter() int64 {
	return c.counter
}

func (c *Country) Subdivisions() map[string]*Subdivision {
	return c.subdivisions
}

func newCountry(name, subdivisionName, webpageName string) *Country {
	country := &Country{
		name:         name,
		subdivisions: make(map[string]*Subdivision),
	}
	country.counter++
	country.subdivisions[subdivisionName] = newSubdivision(subdivisionName, webpageName)
	return country
}

func (c *Country) addToCountry(subdivisionName, webpageName string) {
	c.counter++
	_, found := c.subdivisions[subdivisionName]
	if !found {
		c.subdivisions[subdivisionName] = newSubdivision(subdivisionName, webpageName)
		return
	}
	c.subdivisions[subdivisionName].addToSubdivision(webpageName)
}
