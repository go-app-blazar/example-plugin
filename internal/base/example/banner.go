package example

import (
	"github.com/go-app-blazar/blazar/mix"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type exampleBanner struct {
	app.Compo
	mix.MixHTML
	mix.MixEvents

	IHeadlineText string
}

func Banner() *exampleBanner {
	return &exampleBanner{}
}

func (c *exampleBanner) HeadlineText(text string) *exampleBanner {
	c.IHeadlineText = text
	return c
}

func (c *exampleBanner) Render() app.UI {
	element := app.Div().
		Class("example-banner").
		Body(
			app.Text(c.IHeadlineText),
		)

	var finalElement app.UI = element

	c.MixHTML.Wrap(finalElement)
	c.MixEvents.Wrap(finalElement)

	return finalElement
}
