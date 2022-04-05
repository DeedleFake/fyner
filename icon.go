package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/state"
)

// Icon wraps widget.Icon to provide a component that displays images.
type Icon struct {
	once sync.Once
	w    *widget.Icon

	Resource       state.State[fyne.Resource]
	resourceCancel state.CancelFunc
}

func (icon *Icon) init() {
	icon.once.Do(func() {
		icon.w = widget.NewIcon(state.Get(icon.Resource))
	})
}

func (icon *Icon) CanvasObject() fyne.CanvasObject {
	icon.init()
	return icon.w
}

func (icon *Icon) Bind() {
	icon.init()
	icon.Unbind()

	icon.resourceCancel = icon.Resource.Listen(icon.w.SetResource)
}

func (icon *Icon) Unbind() {
	cancel(&icon.resourceCancel)
}
