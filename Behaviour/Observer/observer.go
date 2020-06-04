package Observer

import (
	"fmt"
)

type IObserver interface {
	update(count int)
	getId() string
}


type Observer struct {
	id string
}

func (o Observer) getId() string {
	return o.id
}

func (o Observer) update(count int) {
	fmt.Println(count)
}
