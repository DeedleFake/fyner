package fstate

import (
	"errors"
	"sync"

	"deedles.dev/state"
	"fyne.io/fyne/v2/data/binding"
)

// Binding represents a generic Fyne data binding.
type Binding[T any] interface {
	binding.DataItem
	Get() (T, error)
	Set(T) error
}

type fromBinding[T any] struct {
	b Binding[T]
}

// FromBinding returns a state that defers to a Fyne data binding.
//
// TODO: Handle errors somehow.
func FromBinding[T any, B Binding[T]](b B) state.MutableState[T] {
	return fromBinding[T]{b: b}
}

func (b fromBinding[T]) Listen(f func(T)) state.CancelFunc {
	lis := binding.NewDataListener(func() {
		f(b.Get())
	})
	b.b.AddListener(lis)

	f(b.Get())

	return func() {
		b.b.RemoveListener(lis)
	}
}

func (b fromBinding[T]) Set(v T) {
	b.b.Set(v)
}

func (b fromBinding[T]) Get() T {
	v, _ := b.b.Get()
	return v
}

type dataItem[T any] struct {
	s state.State[T]
	m sync.Map
}

// ToBinding creates a Binding from a state.State.
func ToBinding[T any](s state.State[T]) Binding[T] {
	return &dataItem[T]{
		s: s,
	}
}

func (item *dataItem[T]) AddListener(lis binding.DataListener) {
	item.m.Store(lis, func(T) {
		lis.DataChanged()
	})
}

func (item *dataItem[T]) RemoveListener(lis binding.DataListener) {
	item.m.Delete(lis)
}

func (item *dataItem[T]) Get() (T, error) {
	return state.Get(item.s), nil
}

func (item *dataItem[T]) Set(v T) error {
	if s, ok := item.s.(state.Setter[T]); ok {
		s.Set(v)
	}
	return nil
}

type dataList[T any, S state.State[T], L ~[]S] struct {
	s state.State[L]
	m sync.Map
}

// ToListBinding creates a binding.DataList from a state.State containing a
// slice of States.
func ToListBinding[T any, S state.State[T], L ~[]S](s state.State[L]) binding.DataList {
	return &dataList[T, S, L]{
		s: s,
	}
}

func (list *dataList[T, S, L]) AddListener(lis binding.DataListener) {
	list.m.Store(lis, func(L) {
		lis.DataChanged()
	})
}

func (list *dataList[T, S, L]) RemoveListener(lis binding.DataListener) {
	list.m.Delete(lis)
}

func (list *dataList[T, S, L]) GetItem(index int) (binding.DataItem, error) {
	s := state.Get(list.s)
	if (index < 0) || (index >= len(s)) {
		return nil, errors.New("index out of range")
	}

	return ToBinding[T](s[index]), nil
}

func (list *dataList[T, S, L]) Length() int {
	return len(state.Get(list.s))
}
