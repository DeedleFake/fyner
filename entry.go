package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/fyner/state"
)

type Entry struct {
	once sync.Once
	w    *widget.Entry

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
	entry.textCancel = entry.Text.Listen(func(v string) {
		entry.w.SetText(v)
	})
	entry.w.OnChanged = func(v string) {
		entry.Text.Set(v)
	}
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
