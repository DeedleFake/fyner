package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/fyner/state"
)

type Label struct {
	once sync.Once
	w    *widget.Label

	Text       state.State[string]
	textCancel state.CancelFunc
}

func (label *Label) init() {
	label.once.Do(func() {
		label.w = widget.NewLabel("")
		label.bind()
	})
}

func (label *Label) bind() {
	if label.Text != nil {
		label.textCancel = label.Text.Listen(func(v string) {
			label.w.SetText(v)
		})
	}
}

func (label *Label) CanvasObject() fyne.CanvasObject {
	label.init()
	return label.w
}

func (label *Label) Bind() {
	label.init()
	label.bind()
}

func (label *Label) Unbind() {
	cancel(&label.textCancel)
}