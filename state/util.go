package state

// ToSliceOfStates returns a state derived from s that wraps each
// element in a Static state. This is useful for, for example, simple
// usages of fyner.List.
func ToSliceOfStates[T any, L ~[]T, S State[L]](s S) State[[]State[T]] {
	return Derived(s, func(v L) []State[T] {
		r := make([]State[T], 0, len(v))
		for _, v := range v {
			r = append(r, Static(v))
		}
		return r
	})
}
