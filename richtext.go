package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/fyner/state"
)

// RichText wraps widget.RichText to provide a component for
// displaying complex text layouts.
type RichText struct {
	once sync.Once
	w    *widget.RichText

	Markdown       state.State[string]
	markdownCancel state.CancelFunc
}

func (rt *RichText) init() {
	rt.once.Do(func() {
		rt.w = widget.NewRichText()
		rt.bind()
	})
}

func (rt *RichText) bind() {
	if rt.Markdown != nil {
		rt.markdownCancel = rt.Markdown.Listen(rt.w.ParseMarkdown)
	}
}

func (rt *RichText) CanvasObject() fyne.CanvasObject {
	rt.init()
	return rt.w
}
