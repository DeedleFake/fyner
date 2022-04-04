package main

import (
	"strings"

	"fyne.io/fyne/v2/app"
	"github.com/DeedleFake/fyner"
	"github.com/DeedleFake/fyner/state"
)

func main() {
	a := app.New()

	horizontal := state.Mutable(false)

	entry := state.Mutable("This is an example.")
	label := state.Derived(entry, func(v string) string {
		return strings.ToUpper(v)
	})

	w := a.NewWindow("Example")
	w.SetContent(fyner.Content(
		&fyner.Center{
			Child: &fyner.Box{
				Horizontal: horizontal,
				Children: []fyner.Component{
					&fyner.Label{
						Text: label,
					},
					&fyner.Entry{
						Text: entry,
					},
					&fyner.Button{
						Text: state.Static("Toggle Direction"),
						OnTapped: func() {
							horizontal.Set(!state.Get[bool](horizontal))
						},
					},
				},
			},
		},
	))

	w.ShowAndRun()
}
