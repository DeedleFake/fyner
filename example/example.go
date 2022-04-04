package main

import (
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/DeedleFake/fyner"
	"github.com/DeedleFake/fyner/state"
)

func main() {
	a := app.New()

	entry := state.Mutable("This is an example.")
	label := state.Derived(entry, func(v string) string {
		return strings.ToUpper(v)
	})

	w := a.NewWindow("Example")
	w.SetContent(
		container.NewCenter(
			container.NewVBox(
				fyner.Content(&fyner.Label{
					Text: label,
				}),
				fyner.Content(&fyner.Entry{
					Text: entry,
				}),
			),
		),
	)

	w.ShowAndRun()
}
