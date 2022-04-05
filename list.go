package fyner

import (
	"reflect"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/fyner/state"
)

type List[E any, C Component] struct {
	once sync.Once
	w    *widget.List

	Items       state.State[[]state.State[E]]
	itemsCancel state.CancelFunc

	Binder func(state.State[E], C)
}

func (list *List[E, C]) init() {
	list.once.Do(func() {
		// TODO: This leaks memory. Unfortunately, it seems like Fyne
		// probably does the same, so there's little that I think that I
		// can do about it.
		components := make(map[fyne.CanvasObject]C)

		// TODO: Move this into Bind().
		list.w = widget.NewListWithData(
			state.ToListBinding[E](list.Items),
			func() fyne.CanvasObject {
				n := deepZero[C]()
				co := n.CanvasObject()
				components[co] = n
				return co
			},
			func(b binding.DataItem, o fyne.CanvasObject) {
				n := components[o]
				n.Unbind()
				list.Binder(state.FromBinding[E](b.(state.Binding[E])), n)
				n.Bind()
			},
		)
	})
}

func (list *List[E, C]) CanvasObject() fyne.CanvasObject {
	list.init()
	return list.w
}

func (list *List[E, C]) Bind() {
	list.init()
	list.Unbind()
}

func (list *List[E, C]) Unbind() {
}

func deepZero[T any]() (v T) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Pointer {
		return v
	}
	return reflect.New(t.Elem()).Interface().(T)
}
