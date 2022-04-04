package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/DeedleFake/fyner/state"
)

// Container wraps fyne.Container to provide a container component.
type Container struct {
	once sync.Once
	w    *fyne.Container

	// Layout is the layout of the children in the container.
	Layout       state.State[fyne.Layout]
	layoutCancel state.CancelFunc

	// Children is the children in the container. They are displayed
	// according to the value of Layout.
	//
	// TODO: Make this stateful? Alternatively, add a special component
	// that can remove a child from a container?
	Children []Component
}

func (c *Container) init() {
	c.once.Do(func() {
		objects := make([]fyne.CanvasObject, 0, len(c.Children))
		for _, child := range c.Children {
			if child == nil {
				continue
			}
			objects = append(objects, child.CanvasObject())
		}
		c.w = container.NewWithoutLayout(objects...)
		c.bind()
	})
}

func (c *Container) bind() {
	if c.Layout != nil {
		c.layoutCancel = c.Layout.Listen(func(v fyne.Layout) {
			c.w.Layout = v
			c.w.Refresh()
		})
	}
}

func (c *Container) CanvasObject() fyne.CanvasObject {
	c.init()
	return c.w
}

//func (c *Container) Bind() {
//	c.init()
//	c.bind()
//}
//
//func (c *Container) Unbind() {
//	cancel(&c.layoutCancel)
//}

// Center is a container with the Center layout. Unlike the Fyne
// version, it only holds a single child component. To replicate
// Fyne's Center's stacking behavior, use a Container with a center
// layout manually.
type Center struct {
	once sync.Once
	c    *Container

	Child Component
}

func (c *Center) init() {
	c.once.Do(func() {
		c.c = &Container{
			Layout:   state.Static(layout.NewCenterLayout()),
			Children: []Component{c.Child},
		}
	})
}

func (c *Center) CanvasObject() fyne.CanvasObject {
	c.init()
	return c.c.CanvasObject()
}

//func (c *Center) Bind() {
//	c.init()
//	c.c.Bind()
//}
//
//func (c *Center) Unbind() {
//	c.c.Unbind()
//}

// Box is a container that displays components in either a row or a
// column. In other words, it wraps both the HBox and VBox layouts.
type Box struct {
	once sync.Once
	c    *Container

	// Horizontal, if true, results in a row rather than a column. If
	// Horizontal is nil, it is treated as though it were false.
	Horizontal state.State[bool]

	Children []Component
}

func (b *Box) init() {
	b.once.Do(func() {
		horizontal := state.Static(layout.NewVBoxLayout())
		if b.Horizontal != nil {
			horizontal = state.Derived(b.Horizontal, func(h bool) fyne.Layout {
				if h {
					return layout.NewHBoxLayout()
				}
				return layout.NewVBoxLayout()
			})
		}

		b.c = &Container{
			Layout:   horizontal,
			Children: b.Children,
		}
	})
}

func (b *Box) CanvasObject() fyne.CanvasObject {
	b.init()
	return b.c.CanvasObject()
}

//func (b *Box) Bind() {
//	b.init()
//	b.c.Bind()
//}
//
//func (b *Box) Unbind() {
//	b.c.Unbind()
//}

type Border struct {
	once sync.Once
	c    *Container

	Top    Component
	Bottom Component
	Left   Component
	Right  Component
	Center Component
}

func (b *Border) init() {
	b.once.Do(func() {
		var top, bottom, left, right fyne.CanvasObject
		if b.Top != nil {
			top = b.Top.CanvasObject()
		}
		if b.Bottom != nil {
			bottom = b.Bottom.CanvasObject()
		}
		if b.Left != nil {
			left = b.Bottom.CanvasObject()
		}
		if b.Right != nil {
			right = b.Bottom.CanvasObject()
		}

		b.c = &Container{
			Layout:   state.Static(layout.NewBorderLayout(top, bottom, left, right)),
			Children: []Component{b.Top, b.Bottom, b.Left, b.Right, b.Center},
		}
	})
}

func (b *Border) CanvasObject() fyne.CanvasObject {
	b.init()
	return b.c.CanvasObject()
}
