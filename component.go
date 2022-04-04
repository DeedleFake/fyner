package fyner

import (
	"fyne.io/fyne/v2"
	"github.com/DeedleFake/fyner/state"
)

// Component is the interface shared by all Fyner components.
type Component interface {
	CanvasObject() fyne.CanvasObject

	//Bind()
	//Unbind()
}

func Content(c Component) fyne.CanvasObject {
	return c.CanvasObject()
}

// cancel calls *f if it isn't nil and then sets it to nil.
func cancel(f *state.CancelFunc) {
	if *f != nil {
		(*f)()
		*f = nil
	}
}
