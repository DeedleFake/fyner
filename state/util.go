package state

// ToSliceOfStates returns a state derived from s that wraps each
// element in a Static state. This is useful for, for example, simple
// usages of fyner.List.
func ToSliceOfStates[T any, S ~[]T](s State[S]) State[[]State[T]] {
	return Derived(s, func(v S) []State[T] {
		r := make([]State[T], 0, len(v))
		for _, v := range v {
			r = append(r, Static(v))
		}
		return r
	})
}
