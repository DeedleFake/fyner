package fyner

import (
	"sync"

	"deedles.dev/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Entry wraps widget.Entry to provide a text entry component.
//
// TODO: Make sure that this API works correctly for strange cases
// like only allowing uppercase letters.
type Entry struct {
	once sync.Once
	w    *widget.Entry

	// Text is the editable text currently in the entry.
	Text       state.MutableState[string]
	textCancel state.CancelFunc

	// Disabled, if true, prevents user input to the text field.
	Disabled       state.State[bool]
	disabledCancel state.CancelFunc
}

func (entry *Entry) init() {
	entry.once.Do(func() {
		entry.w = widget.NewEntry()
	})
}

func (entry *Entry) CanvasObject() fyne.CanvasObject {
	entry.init()
	return entry.w
}

func (entry *Entry) Bind() {
	entry.init()
	entry.Unbind()

	if entry.Text != nil {
		entry.textCancel = entry.Text.Listen(entry.w.SetText)
		entry.w.OnChanged = entry.Text.Set
	}

	if entry.Disabled != nil {
		entry.disabledCancel = entry.Disabled.Listen(func(disabled bool) {
			if disabled {
				entry.w.Disable()
				return
			}
			entry.w.Enable()
		})
	}
}

func (entry *Entry) Unbind() {
	cancel(&entry.textCancel)
	entry.w.OnChanged = nil

	cancel(&entry.disabledCancel)
}
