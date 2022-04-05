// Package fyner provides a declarative wrapper around the Fyne UI
// library.
//
// Fyner's approach to structuring the UI and state management of an
// app is very difference from Fyne's. Fyner provides its own set of
// widgets that wrap the basic Fyne widgets, providing similar
// functionality but in a far more declarative way. To differentiate,
// Fyner refers to its own widgets as "components".
//
// Fyner components are instantiated manually as struct pointer
// literals. They export a number of fields which may be set, some of
// which are of a type from the state package. The fields of a
// component should not be changed after it is created, though if any
// of the fields are mutable state types, they may be set via the
// state API.
//
// For example, to create a center-layout container that contains a
// single label:
//
//    &fyner.Center{
//    	Child: &fyner.Label{
//    		Text: state.Static("This is an example."),
//    	},
//    }
//
// To interact with the typical Fyne UI system, a Content function is
// provided that turns a Fyner component into a Fyne CanvasObject.
// This is typically only used at the top-level in order to set the
// content of a window, hence the name, but it can actually be used
// anywhere that a client might want to insert a component into a
// regular Fyne UI layout.
//
// To illustrate the whole system, here's a complete example:
//
//    package main
//
//    import (
//    	"deedles.dev/fyner"
//    	"deedles.dev/fyner/state"
//    	"fyne.io/fyne/v2/app"
//    )
//
//    func main() {
//    	a := app.New()
//
//    	text := state.Mutable("This is an example.")
//
//    	w := a.NewWindow("Example")
//    	w.SetContent(fyner.Content(
//    		&fyner.Center{
//    			Child: &fyner.Box{
//    				Children: []fyner.Component{
//    					&fyner.Label{
//    						Text: text,
//    					},
//    					&fyner.Button{
//    						Text: state.Static("Greet"),
//    						OnTapped: func() {
//    							text.Set("Hi.")
//    						},
//    					},
//    				},
//    			},
//    		},
//    	))
//
//    	w.ShowAndRun()
//    }
package fyner
