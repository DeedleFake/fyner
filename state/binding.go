package state

import (
	fyneBinding "fyne.io/fyne/v2/data/binding"
)

// Binding represents a generic Fyne data binding.
type Binding[T any] interface {
	fyneBinding.DataItem
	Get() (T, error)
	Set(T) error
}

type binding[T any] struct {
	b Binding[T]
}

// FromBinding returns a state that defers to a Fyne data binding.
//
// TODO: Handle errors somehow.
func FromBinding[T any, B Binding[T]](b B) MutableState[T] {
	return binding[T]{b: b}
}

func (b binding[T]) Listen(f func(T)) CancelFunc {
	lis := fyneBinding.NewDataListener(func() {
		f(b.Get())
	})
	b.b.AddListener(lis)
	return func() {
		b.b.RemoveListener(lis)
	}
}

func (b binding[T]) Set(v T) {
	b.b.Set(v)
}

func (b binding[T]) Get() T {
	v, _ := b.b.Get()
	return v
}
