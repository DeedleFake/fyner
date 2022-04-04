package state

import "sync"

// State represents a value that can change over time.
type State[T any] interface {
	// Listen registers a listener function to be called whenever the
	// State's value changes.
	//
	// When Listen is first called, it immedietely calls the provided
	// function with the current value of the state and does not return
	// until after the provided function returns.
	Listen(func(T)) CancelFunc
}

// Getter is an optional interface defined by States to allow for
// quicker one-time fetching of their current value.
type Getter[T any] interface {
	Get() T
}

// Setter represents a value that can be changed. Setting a value is
// thread-safe with reading it.
type Setter[T any] interface {
	Set(v T)
}

// MutableState is a State that can be changed. There is no guarantee
// about the relationship between the Set method returning and the
// State's listeners being informed of the update.
type MutableState[T any] interface {
	State[T]
	Setter[T]
}

// Get returns the current value of a given state.
func Get[T any](s State[T]) (v T) {
	if g, ok := s.(Getter[T]); ok {
		return g.Get()
	}

	var cancel CancelFunc
	cancel = s.Listen(func(c T) {
		defer cancel()
		v = c
	})
	return v
}

type CancelFunc func()

type static[T any] struct {
	v T
}

// Static returns a state that never changes. If a value is definitely
// constant, using this type of state is more efficient than a mutable
// one.
func Static[T any](v T) State[T] {
	return static[T]{v: v}
}

func (s static[T]) Listen(listener func(T)) CancelFunc {
	listener(s.v)
	return func() {}
}

func (s static[T]) Get() T {
	return s.v
}

type mutable[T any] struct {
	lis Listenable[T]
	m   sync.RWMutex
	v   T
}

func Mutable[T any](v T) MutableState[T] {
	return &mutable[T]{v: v}
}

func (s *mutable[T]) Listen(f func(T)) CancelFunc {
	id := s.lis.Add(f)
	return func() {
		s.lis.Remove(id)
	}
}

func (s *mutable[T]) Set(v T) {
	s.m.Lock()
	defer s.m.Unlock()

	s.v = v
	s.lis.Send(v)
}

func (s *mutable[T]) Get() T {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.v
}

type derived[T, F any] struct {
	from State[F]
	m    func(F) T
}

// Derived returns a read-only state that derives its values from
// another state, passing them through the mapping function m. In
// other words, when from's value changes, the derived state's
// listeners will be called with that new value passed through m.
func Derived[T, F any](from State[F], m func(F) T) State[T] {
	return derived[T, F]{
		from: from,
		m:    m,
	}
}

func (s derived[T, F]) Listen(f func(T)) CancelFunc {
	return s.from.Listen(func(v F) {
		f(s.m(v))
	})
}

func (s derived[T, F]) Get() T {
	return s.m(Get(s.from))
}
