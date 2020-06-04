package Decorator

import "fmt"

func Client(c IComponent) {
	fmt.Println(c.print())
}

func Run() {
	a := &component{
		name:  "Box",
		count: 0,
	}
	b := NewDecoratorA(a)
	c := NewDecoratorB(b)
	Client(a)
	Client(b)
	Client(c)
}
