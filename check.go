package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/DeedleFake/fyner/state"
)

// Check wraps widget.Check to provide a toggleable checkbox
// component.
type Check struct {
	once sync.Once
	w    *widget.Check

	Text       state.State[string]
	textCancel state.CancelFunc

	Checked       state.MutableState[bool]
	checkedCancel state.CancelFunc
}

func (check *Check) init() {
	check.once.Do(func() {
		check.w = widget.NewCheck("", nil)
		check.bind()
	})
}

func (check *Check) bind() {
	if check.Text != nil {
		check.textCancel = check.Text.Listen(func(v string) {
			check.w.Text = v
			check.w.Refresh()
		})
	}

	if check.Checked != nil {
		check.checkedCancel = check.Checked.Listen(check.w.SetChecked)
		check.w.OnChanged = check.Checked.Set
	}
}

func (check *Check) CanvasObject() fyne.CanvasObject {
	check.init()
	return check.w
}