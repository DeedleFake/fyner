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

	// Markdown provides markdown source as a string to generate the
	// rich text from. It is not recommended to combine this with manual
	// text segments, as any changes to either one will override the
	// latest value of the other.
	Markdown       state.State[string]
	markdownCancel state.CancelFunc

	// Segments is the list of RichTextSegments to display with the
	// component.
	Segments       state.State[[]widget.RichTextSegment]
	segmentsCancel state.CancelFunc
}

func (rt *RichText) init() {
	rt.once.Do(func() {
		rt.w = widget.NewRichText()
	})
}

func (rt *RichText) CanvasObject() fyne.CanvasObject {
	rt.init()
	return rt.w
}

func (rt *RichText) Bind() {
	rt.init()
	rt.Unbind()

	if rt.Markdown != nil {
		rt.markdownCancel = rt.Markdown.Listen(rt.w.ParseMarkdown)
	}

	if rt.Segments != nil {
		rt.segmentsCancel = rt.Segments.Listen(func(s []widget.RichTextSegment) {
			rt.w.Segments = s
			rt.w.Refresh()
		})
	}
}

func (rt *RichText) Unbind() {
	cancel(&rt.markdownCancel)
}
