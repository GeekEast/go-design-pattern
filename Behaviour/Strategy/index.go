package strategy

func Run() {
	strategy1 := algorithm1Entity{}
	strategy2 := algorithm2Entity{}
	c := NewContext(&strategy1)
	c.execute()
	c.setStrategy(&strategy2)
	c.execute()
}
