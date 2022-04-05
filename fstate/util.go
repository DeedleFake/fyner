package fstate

import "deedles.dev/state"

// ToSliceOfStates returns a state derived from s that wraps each
// element in a state.Static. This is useful for, for example, simple
// usages of fyner.List.
func ToSliceOfStates[T any, L ~[]T, S state.State[L]](s S) state.State[[]state.State[T]] {
	return state.Derived(s, func(v L) []state.State[T] {
		r := make([]state.State[T], 0, len(v))
		for _, v := range v {
			r = append(r, state.Static(v))
		}
		return r
	})
}
