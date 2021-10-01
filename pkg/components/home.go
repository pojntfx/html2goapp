package components

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/html-to-go-app-converter/pkg/converter"
)

type Home struct {
	app.Compo

	input string

	goAppPkg  string
	pkg       string
	component string

	output string
}

func (c *Home) Render() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.H1().Text("HTML to go-app Converter"),
		),
		app.Main().Body(
			app.Section().ID("input-section").Body(
				app.Form().OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					generated, err := converter.ConvertHTMLToComponent(
						c.input,
						c.goAppPkg,
						c.pkg,
						c.component,
					)
					if err != nil {
						generated += err.Error()

						log.Println("could not convert HTML to component:", err)
					}

					c.output = generated
				}).Body(
					app.Label().Text("HTML Input").For("html-input"),
					app.Br(),
					app.Textarea().
						ID("html-input").
						Placeholder("Enter HTML here").
						Required(true).
						OnInput(func(ctx app.Context, e app.Event) {
							c.input = ctx.JSSrc().Get("value").String()
						}).
						Style("width", "100%").
						Style("resize", "vertical").
						Rows(25).
						Text(c.input),

					app.Br(),
					app.Label().Text("go-app Package").For("go-app-pkg-input"),
					app.Br(),
					app.Input().
						ID("go-app-pkg-input").
						Required(true).
						OnInput(func(ctx app.Context, e app.Event) {
							c.goAppPkg = ctx.JSSrc().Get("value").String()
						}).
						Value(c.goAppPkg),

					app.Br(),
					app.Label().Text("Component Package").For("component-pkg-input"),
					app.Br(),
					app.Input().
						ID("component-pkg-input").
						Required(true).
						OnInput(func(ctx app.Context, e app.Event) {
							c.pkg = ctx.JSSrc().Get("value").String()
						}).
						Value(c.pkg),

					app.Br(),
					app.Label().Text("Component Name").For("component-name-input"),
					app.Br(),
					app.Input().
						ID("component-name-input").
						Required(true).
						OnInput(func(ctx app.Context, e app.Event) {
							c.component = ctx.JSSrc().Get("value").String()
						}).
						Value(c.component),

					app.Br(),
					app.Button().
						Text("Convert").
						Type("submit"),
				),
			),
			app.Section().ID("output-section").Body(
				app.Label().Text("go-app Output").For("go-app-output"),
				app.Br(),
				app.Textarea().
					ID("go-app-output").
					Placeholder("go-app's syntax will be here").
					ReadOnly(true).
					Style("width", "100%").
					Style("resize", "vertical").
					Rows(25).
					Text(c.output),
			),
		),
		app.Footer().Body(
			app.A().Href("https://github.com/pojntfx/html2goapp").Target("_blank").Text("© 2021 AGPL-3.0 Felicitas Pojtinger"),
		),
	)
}

func (c *Home) OnMount(app.Context) {
	c.goAppPkg = "github.com/maxence-charriere/go-app/v9/pkg/app"
	c.pkg = "components"
	c.component = "MyComponent"
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
