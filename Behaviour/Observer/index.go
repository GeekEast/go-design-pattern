package Observer

import "time"

func Client() {
	store := NewStore()
	consumer1 := Observer{"1"}
	consumer2 := Observer{"2"}
	store.add(consumer1)
	store.add(consumer2)

	store.SetState(1)

	time.Sleep(2 * time.Second)
	store.SetState(2)

	time.Sleep(2 * time.Second)
	store.SetState(3)
}
