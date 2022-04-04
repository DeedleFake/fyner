package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/fyner/state"
)

type List[E any] struct {
	once sync.Once
	w    *widget.List

	Items       state.State[[]state.State[E]]
	itemsCancel state.CancelFunc

	ItemBuilder func() Component
}

func (list *List[E]) init() {
	list.once.Do(func() {
		//components := make(map[int]Component)

		list.w = widget.NewListWithData(
			state.ToListBinding[E](list.Items),
			func() fyne.CanvasObject { panic("Not implemented.") },
			func(b binding.DataItem, o fyne.CanvasObject) {
				panic("Not implemented.")
			},
		)
	})
}

func (list *List[E]) CanvasObject() fyne.CanvasObject {
	list.init()
	return list.w
}
