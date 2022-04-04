package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/fyner/state"
)

// Button wraps widget.Button to provide a button component.
type Button struct {
	once sync.Once
	w    *widget.Button

	// Text is the text label displayed on the button.
	Text       state.State[string]
	textCancel state.CancelFunc

	// OnTapped is called when the button is tapped/clicked on.
	OnTapped func()
}

func (button *Button) init() {
	button.once.Do(func() {
		button.w = widget.NewButton("", button.OnTapped)
		button.bind()
	})
}

func (button *Button) bind() {
	if button.Text != nil {
		button.textCancel = button.Text.Listen(func(v string) {
			button.w.SetText(v)
		})
	}
}

func (button *Button) CanvasObject() fyne.CanvasObject {
	button.init()
	return button.w
}

//func (button *Button) Bind() {
//	button.init()
//	button.bind()
//}
//
//func (button *Button) Unbind() {
//	cancel(&button.textCancel)
//}
