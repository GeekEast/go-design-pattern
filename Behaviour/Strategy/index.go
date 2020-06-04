package strategy

import (
	"errors"
	"fmt"
)

// Definition of Strategy
type Strategy interface {
	algorithm()
}

// strategy 1
type algorithm1Entity struct{}

func (e *algorithm1Entity) algorithm() {
	fmt.Println("algorithm 1")
}

// strategy 2
type algorithm2Entity struct{}

func (e *algorithm2Entity) algorithm() {
	fmt.Println("algorithm 2")
}

// Retrieve Strategy
// get strategy from global map
var Store = map[string]Strategy{
	"behaviour1": &algorithm1Entity{},
	"behaviour2": &algorithm2Entity{},
}

// get strategy from generator function (when you have specific state for each use case)
func getStrategy(t string) (s Strategy, err error) {
	if t == "behaviour1" {
		return &algorithm1Entity{}, nil
	} else if t == "behaviour2" {
		return &algorithm2Entity{}, nil
	} else {
		return nil, errors.New("no such strategy")
	}
}

// Use strategy
func use(t string) {
	// stateless
	Store[t].algorithm()

	// stateful
	behaviourEntity, err := getStrategy(t)
	if err != nil {
		fmt.Println(err)
	}
	behaviourEntity.algorithm()
}