package Observer

// design the interface at first
// interface help decouple the Publisher and the Subscriber
type Observable interface {
	add(o IObserver)
	remove(o IObserver)
	notifyAll()
}

// struct is just a class
type Store struct {
	ObserverList []IObserver
	count        int
}

// design your own constructor
func NewStore() *Store {
	return &Store{
		ObserverList: []IObserver{},
		count:        0,
	}
}

// implement Observable interface
func (s *Store) notifyAll() {
	for _, observer := range s.ObserverList {
		observer.update(s.count)
	}
}

func (s *Store) add(o IObserver) {
	s.ObserverList = append(s.ObserverList, o)
}

func (s *Store) remove(o IObserver) {
	s.ObserverList = removeFromSlice(s.ObserverList, o)
}

// extra methods
func (s *Store) SetState(c int) {
	s.count = c
	s.notifyAll()
}

func (s *Store) GetState() int {
	return s.count
}
