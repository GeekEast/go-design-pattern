package strategy


type context struct {
	strategy IStrategy
}

func (c *context) execute() {
	c.strategy.algorithm()
}

func (c *context) setStrategy(s IStrategy) {
	c.strategy = s
}

// constructor
func NewContext(s IStrategy) *context {
	return &context{strategy: s}
}
