package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Home struct {
	app.Compo

	input  string
	output string
}

func (c *Home) Render() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.H1().Text("HTML to go-app Converter"),
		),
		app.Main().Body(
			app.Section().ID("input").Body(
				app.Textarea().
					Placeholder("Enter HTML here").
					OnInput(func(ctx app.Context, e app.Event) {
						c.input = ctx.JSSrc().Get("value").String()
					}).
					Text(c.input),
			),
			app.Section().ID("actions").Body(
				app.Button().
					Text("Convert").
					OnClick(func(ctx app.Context, e app.Event) {
						c.output = c.input
					}),
			),
			app.Section().ID("output").Body(
				app.Textarea().
					Placeholder("go-app's syntax will be here").
					ReadOnly(true).
					Text(c.output),
			),
		),
		app.Footer().Body(
			app.A().Href("https://github.com/pojntfx/html2goapp").Target("_blank").Text("© 2021 AGPL-3.0 Felicitas Pojtinger"),
		),
	)
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
