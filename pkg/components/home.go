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
			app.Form().
				Class("pf-c-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
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
				}).
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("go-app-pkg-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("go-App Package"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Required(true).
										OnInput(func(ctx app.Context, e app.Event) {
											c.goAppPkg = ctx.JSSrc().Get("value").String()
										}).
										Value(c.goAppPkg).
										Type("text").
										ID("go-app-pkg-input"),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("component-pkg-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Target Package"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Required(true).
										OnInput(func(ctx app.Context, e app.Event) {
											c.pkg = ctx.JSSrc().Get("value").String()
										}).
										Value(c.pkg).
										Type("text").
										ID("component-pkg-input"),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("component-name-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Component Name"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Type("text").
										Required(true).
										OnInput(func(ctx app.Context, e app.Event) {
											c.component = ctx.JSSrc().Get("value").String()
										}).
										Value(c.component).
										ID("component-name-input"),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("html-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Source Code"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Div().
										Class("pf-c-code-editor").
										Body(
											app.Div().
												Class("pf-c-code-editor__header").
												Body(
													app.Div().
														Class("pf-c-code-editor__tab").
														Body(
															app.Span().
																Class("pf-c-code-editor__tab-icon").
																Body(
																	app.I().Class("fas fa-code"),
																),
															app.Span().
																Class("pf-c-code-editor__tab-text").
																Text("HTML"),
														),
												),
											app.Div().
												Class("pf-c-code-editor__main").
												Body(
													app.Textarea().
														ID("html-input").
														Placeholder("Enter HTML input here").
														Required(true).
														OnInput(func(ctx app.Context, e app.Event) {
															c.input = ctx.JSSrc().Get("value").String()
														}).
														Style("width", "100%").
														Style("resize", "vertical").
														Style("border", "0").
														Class("pf-c-form-control").
														Rows(25).
														Text(c.input),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Div().
										Class("pf-c-form__actions").
										Body(
											app.Button().
												Class("pf-c-button pf-m-primary").
												Type("submit").
												Text("Convert to Go"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-code-editor pf-m-read-only").
				Body(
					app.Div().
						Class("pf-c-code-editor__header").
						Body(
							app.Div().
								Class("pf-c-code-editor__tab").
								Body(
									app.Span().
										Class("pf-c-code-editor__tab-icon").
										Body(
											app.I().Class("fas fa-code"),
										),
									app.Span().
										Class("pf-c-code-editor__tab-text").Text("Go"),
								),
						),
					app.Div().
						Class("pf-c-code-editor__main").
						Body(
							app.Textarea().
								Placeholder("go-app's syntax will be here").
								ReadOnly(true).
								Style("width", "100%").
								Style("resize", "vertical").
								Style("border", "0").
								Class("pf-c-form-control").
								Rows(25).
								Text(c.output),
						),
				),
		),
		app.Footer().Body(
			app.A().Href("https://github.com/pojntfx/html2goapp").Target("_blank").Text("© 2021 AGPL-3.0 Felix Pojtinger"),
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
