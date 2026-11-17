package example

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *exampleBanner) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *exampleBanner {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *exampleBanner) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *exampleBanner {
	c.MixOnClick(function)
	return c
}

func (c *exampleBanner) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *exampleBanner {
	c.MixOnDoubleClick(function)
	return c
}

func (c *exampleBanner) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *exampleBanner {
	c.MixOnMouseUp(function)
	return c
}

func (c *exampleBanner) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *exampleBanner {
	c.MixOnMouseDown(function)
	return c
}

func (c *exampleBanner) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *exampleBanner {
	c.MixOnKeyDown(function)
	return c
}

func (c *exampleBanner) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *exampleBanner {
	c.MixOnKeyPress(function)
	return c
}

func (c *exampleBanner) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *exampleBanner {
	c.MixOnKeyUp(function)
	return c
}
