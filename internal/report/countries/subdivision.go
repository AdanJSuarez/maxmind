package countries

type Subdivision struct {
	name     string
	counter  int64
	webpages map[string]int64
}

func (s *Subdivision) Name() string {
	return s.name
}

func (s *Subdivision) Counter() int64 {
	return s.counter
}

func (s *Subdivision) Webpages() map[string]int64 {
	return s.webpages
}

func newSubdivision(name, webpageName string) *Subdivision {
	subdivision := &Subdivision{
		name:     name,
		counter:  1,
		webpages: make(map[string]int64),
	}
	subdivision.webpages[webpageName]++
	return subdivision
}

func (s *Subdivision) addToSubdivision(webpageName string) {
	s.counter++
	s.webpages[webpageName]++
}
