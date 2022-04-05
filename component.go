package fyner

import (
	"fyne.io/fyne/v2"
	"github.com/DeedleFake/state"
)

// Component is the interface shared by all Fyner components.
type Component interface {
	CanvasObject() fyne.CanvasObject

	Bind()
	Unbind()
}

// Content binds c to its state and then returns a CanvasObject that
// can be used to place c into a standard Fyne widget tree.
func Content(c Component) fyne.CanvasObject {
	c.Bind()
	return c.CanvasObject()
}

// cancel calls *f if it isn't nil and then sets it to nil.
func cancel(f *state.CancelFunc) {
	if *f != nil {
		(*f)()
		*f = nil
	}
}
