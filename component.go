package fyner

import (
	"fyne.io/fyne/v2"
	"github.com/DeedleFake/fyner/state"
)

type Component interface {
	CanvasObject() fyne.CanvasObject

	//Bind()
	//Unbind()
}

func Content(c Component) fyne.CanvasObject {
	return c.CanvasObject()
}

func cancel(f *state.CancelFunc) {
	if *f != nil {
		(*f)()
		*f = nil
	}
}
