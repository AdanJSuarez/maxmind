package countries

type Info struct {
	name    string
	visit   int64
	topPage string
}

func NewInfo(name, topPage string, visit int64) Info {
	return Info{
		name:    name,
		topPage: topPage,
		visit:   visit,
	}
}

func (i Info) Name() string {
	return i.name
}

func (i Info) Visit() int64 {
	return i.visit
}

func (i Info) TopPage() string {
	return i.topPage
}
