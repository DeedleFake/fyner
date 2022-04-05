package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/state"
)

// Button wraps widget.Button to provide a button component.
type Button struct {
	once sync.Once
	w    *widget.Button

	// Text is the text label displayed on the button.
	Text       state.State[string]
	textCancel state.CancelFunc

	// Disabled, if true, disables the button, prevening input.
	Disabled       state.State[bool]
	disabledCancel state.CancelFunc

	// OnTapped is called when the button is tapped/clicked on.
	OnTapped func()
}

func (button *Button) init() {
	button.once.Do(func() {
		button.w = widget.NewButton("", button.OnTapped)
	})
}

func (button *Button) CanvasObject() fyne.CanvasObject {
	button.init()
	return button.w
}

func (button *Button) Bind() {
	button.init()
	button.Unbind()

	if button.Text != nil {
		button.textCancel = button.Text.Listen(button.w.SetText)
	}

	if button.Disabled != nil {
		button.disabledCancel = button.Disabled.Listen(func(disabled bool) {
			if disabled {
				button.w.Disable()
				return
			}
			button.w.Enable()
		})
	}
}

func (button *Button) Unbind() {
	cancel(&button.textCancel)
}
