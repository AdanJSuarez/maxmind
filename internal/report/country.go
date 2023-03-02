package report

type country struct {
	counter  int64
	webpages map[string]int64
}

func newCountry(webpageName string) *country {
	c := &country{
		counter:  0,
		webpages: make(map[string]int64),
	}
	c.add(webpageName)
	return c
}

func (c *country) add(webpageName string) {
	c.counter++
	_, found := c.webpages[webpageName]
	if !found {
		c.webpages[webpageName] = 1
		return
	}
	c.webpages[webpageName]++
}
