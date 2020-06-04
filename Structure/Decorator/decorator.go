package Decorator

type decoratorBase struct {
	c IComponent
}

func (d decoratorBase) getCount() int {
	return 1
}
func (d decoratorBase) print() string {
	return "decorator base " + d.c.print()
}

type decoratorForPrintA struct {
	decoratorBase
}

func (d decoratorForPrintA) print() string {
	return "A " + d.decoratorBase.print()
}

func NewDecoratorA(c IComponent) IComponent {
	return &decoratorForPrintA{decoratorBase{c}}
}

type decoratorForPrintB struct {
	decoratorBase
}

func (d decoratorForPrintB) print() string {
	return "B" + d.decoratorBase.print()
}

func NewDecoratorB(c IComponent) IComponent {
	return &decoratorForPrintB{decoratorBase{c}}
}
