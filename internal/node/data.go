package node

// Data represents the single instance of data.
type Data struct {
	name    string
	counter int64
}

// NewData returns an initialized instance of Data.
func NewData(name string, counter int64) Data {
	return Data{
		name:    name,
		counter: counter,
	}
}

// Getters

func (d *Data) Name() string {
	return d.name
}

func (d *Data) Counter() int64 {
	return d.counter
}
