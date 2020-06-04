package Decorator


type component struct {
	name string
	count int
}

func (c component) getCount() int {
	return c.count
}

func (c component) print() string {
	return "My Name is " + c.name
}
