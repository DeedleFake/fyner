package fyner

import (
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

	Builder func() C
	Binder  func(state.State[E], C)
}

func (list *List[E, C]) init() {
	list.once.Do(func() {
		unbound := make(map[fyne.CanvasObject]C)
		bound := make(map[fyne.CanvasObject]C)

		// TODO: Move this into Bind().
		list.w = widget.NewListWithData(
			state.ToListBinding[E](list.Items),
			func() fyne.CanvasObject {
				n := list.Builder()
				co := n.CanvasObject()
				unbound[co] = n
				return co
			},
			func(b binding.DataItem, o fyne.CanvasObject) {
				n := unbound[o]
				delete(unbound, o)
				list.Binder(state.FromBinding[E](b.(state.Binding[E])), n)
				bound[o] = n
			},
		)
	})
}

func (list *List[E, C]) CanvasObject() fyne.CanvasObject {
	list.init()
	return list.w
}

func (list *List[E, C]) Bind() {
}

func (list *List[E, C]) Unbind() {
}
