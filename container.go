package fyner

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/DeedleFake/fyner/state"
)

type Container struct {
	once sync.Once
	w    *fyne.Container

	Layout       state.State[fyne.Layout]
	layoutCancel state.CancelFunc

	Children []Component
}

func (c *Container) init() {
	c.once.Do(func() {
		objects := make([]fyne.CanvasObject, 0, len(c.Children))
		for _, child := range c.Children {
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

func (c *Container) Bind() {
	c.init()
	c.bind()
}

func (c *Container) Unbind() {
	cancel(&c.layoutCancel)
}

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

func (c *Center) Bind() {
	c.init()
	c.c.Bind()
}

func (c *Center) Unbind() {
	c.c.Unbind()
}

type Box struct {
	once sync.Once
	c    *Container

	Horizontal state.State[bool]

	Children []Component
}

func (b *Box) init() {
	b.once.Do(func() {
		b.c = &Container{
			Layout: state.Derived(b.Horizontal, func(h bool) fyne.Layout {
				if h {
					return layout.NewHBoxLayout()
				}
				return layout.NewVBoxLayout()
			}),
			Children: b.Children,
		}
	})
}

func (b *Box) CanvasObject() fyne.CanvasObject {
	b.init()
	return b.c.CanvasObject()
}

func (b *Box) Bind() {
	b.init()
	b.c.Bind()
}

func (b *Box) Unbind() {
	b.c.Unbind()
}
