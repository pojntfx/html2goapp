package components

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/html2goapp/pkg/converter"
	"github.com/yosssi/gohtml"
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
	return app.Div().
		Class("pf-c-page").
		Body(
			app.A().
				Class("pf-c-skip-to-content pf-c-button pf-m-primary").
				Href("#main").
				Text("Skip to content"),
			app.Header().
				Class("pf-c-page__header").
				Body(
					app.Div().
						Class("pf-c-page__header-brand").
						Body(
							app.A().
								Href("#").
								Class("pf-c-page__header-brand-link").
								Body(
									app.Img().
										Class("pf-c-brand").
										Src("/web/logo.png").
										Alt("Logo"),
								),
						),
					app.Div().
						Class("pf-c-page__header-tools").
						Body(
							app.Div().
								Class("pf-c-page__header-tools-group").
								Body(
									app.Div().
										Class("pf-c-page__header-tools-item").
										Body(
											app.A().
												Href("https://github.com/pojntfx/html2goapp").
												Target("_blank").
												Class("pf-c-button pf-m-plain").
												Aria("label", "Help").
												Body(
													app.I().
														Class("pf-icon pf-icon-help").
														Aria("hidden", true),
												),
										),
								),
						),
				),
			app.Main().
				ID("main").
				Class("pf-c-page__main").
				TabIndex(-1).
				Body(app.Section().
					Class("pf-c-page__main-section pf-m-fill").
					Body(
						app.Div().
							Class("pf-l-grid pf-m-gutter pf-m-all-6-col-on-xl").
							Body(
								app.Div().
									Class("pf-l-grid__item").
									Body(
										app.Div().
											Class("pf-c-card").
											Body(
												app.Div().
													Class("pf-c-card__title").
													Text("Input"),
												app.Div().
													Class("pf-c-card__body").
													Body(
														app.Form().
															Class("pf-c-form").
															OnSubmit(func(ctx app.Context, e app.Event) {
																e.PreventDefault()

																c.convert()
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
																									Class("pf-c-code-editor__controls").
																									Body(
																										app.Button().
																											Class("pf-c-button pf-m-control").
																											Type("button").
																											Aria("label", "Format").
																											OnClick(func(ctx app.Context, e app.Event) {
																												c.input = gohtml.Format(c.input)
																											}).
																											Body(
																												app.I().
																													Class("fas fa-magic").
																													Aria("hidden", true),
																											),
																									),
																								app.Div().
																									Class("pf-c-code-editor__tab").
																									Body(
																										app.Span().
																											Class("pf-c-code-editor__tab-icon").
																											Body(
																												app.I().
																													Class("fas fa-code"),
																											),
																										app.Span().
																											Class("pf-c-code-editor__tab-text").
																											Body(
																												app.Text("HTML"),
																											),
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
																										if c.input == "" {
																											c.output = ""

																											return
																										}

																										c.convert()
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
													),
											),
									),
								app.Div().
									Class("pf-l-grid__item").
									Body(
										app.Div().
											Class("pf-c-card").
											Body(
												app.Div().
													Class("pf-c-card__title").
													Text("Output"),
												app.Div().
													Class("pf-c-card__body").
													Body(
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
											),
									),
							),
					),
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

func (c *Home) convert() {
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
}
