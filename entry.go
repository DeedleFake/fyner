package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/fyner/state"
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
}

func (entry *Entry) init() {
	entry.once.Do(func() {
		entry.w = widget.NewEntry()
		entry.bind()
	})
}

func (entry *Entry) bind() {
	entry.textCancel = entry.Text.Listen(entry.w.SetText)
	entry.w.OnChanged = entry.Text.Set
}

func (entry *Entry) CanvasObject() fyne.CanvasObject {
	entry.init()
	return entry.w
}

//func (entry *Entry) Bind() {
//	entry.init()
//	entry.bind()
//}
//
//func (entry *Entry) Unbind() {
//	cancel(&entry.textCancel)
//}
