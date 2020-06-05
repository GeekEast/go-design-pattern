package strategy

import "fmt"

type IStrategy interface {
	algorithm()
}

type algorithm1Entity struct{}
func (e *algorithm1Entity) algorithm() {
	fmt.Println("algorithm 1")
}

type algorithm2Entity struct{}
func (e *algorithm2Entity) algorithm() {
	fmt.Println("algorithm 2")
}