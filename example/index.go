package example

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PF4Tabs struct {
	app.Compo
}

func (c *PF4Tabs) Render() app.UI {
	return app.Div().
		Body(
			app.Div().
				Class("pf-c-tabs pf-m-fill").
				DataSet("ouia-component-type", "PF4/Tabs").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-Tabs-12").
				Body(
					app.Button().
						Class("pf-c-tabs__scroll-button").
						Aria("label", "Scroll left").
						Aria("hidden", true).
						Disabled(true).
						Body(),
					app.Ul().
						Class("pf-c-tabs__list").
						Body(
							app.Li().
								Class("pf-c-tabs__item pf-m-current").
								Body(
									app.Button().
										DataSet("ouia-component-type", "PF4/TabButton").
										DataSet("ouia-safe", true).
										Class("pf-c-tabs__link").
										ID("pf-tab-0-pf-1633041198065oegccuf7ng").
										Aria("controls", "pf-tab-section-0-pf-1633041198065oegccuf7ng").
										Body(
											app.Span().
												Class("pf-c-tabs__item-icon"),
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Users"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										DataSet("ouia-component-type", "PF4/TabButton").
										DataSet("ouia-safe", true).
										Class("pf-c-tabs__link").
										ID("pf-tab-1-pf-1633041198065oegccuf7ng").
										Aria("controls", "pf-tab-section-1-pf-1633041198065oegccuf7ng").
										Body(
											app.Span().
												Class("pf-c-tabs__item-icon"),
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Containers"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										DataSet("ouia-component-type", "PF4/TabButton").
										DataSet("ouia-safe", true).
										Class("pf-c-tabs__link").
										ID("pf-tab-2-pf-1633041198065oegccuf7ng").
										Aria("controls", "pf-tab-section-2-pf-1633041198065oegccuf7ng").
										Body(
											app.Span().
												Class("pf-c-tabs__item-icon"),
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Database"),
												),
										),
								),
						),
					app.Button().
						Class("pf-c-tabs__scroll-button").
						Aria("label", "Scroll right").
						Aria("hidden", true).
						Disabled(true).
						Body(),
				),
			app.Section().
				Class("pf-c-tab-content").
				ID("pf-tab-section-0-pf-1633041198065oegccuf7ng").
				Aria("labelledby", "pf-tab-0-pf-1633041198065oegccuf7ng").
				Aria("role", "tabpanel").
				TabIndex(0).
				DataSet("ouia-component-type", "PF4/TabContent").
				DataSet("ouia-safe", true).
				Body(
					app.Text("\n    Users\n  "),
				),
			app.Section().
				Class("pf-c-tab-content").
				ID("pf-tab-section-1-pf-1633041198065oegccuf7ng").
				Aria("labelledby", "pf-tab-1-pf-1633041198065oegccuf7ng").
				Aria("role", "tabpanel").
				TabIndex(0).
				DataSet("ouia-component-type", "PF4/TabContent").
				DataSet("ouia-safe", true).
				Hidden(true).
				Body(
					app.Text("\n    Containers\n  "),
				),
			app.Section().
				Hidden(true).
				Class("pf-c-tab-content").
				ID("pf-tab-section-2-pf-1633041198065oegccuf7ng").
				Aria("labelledby", "pf-tab-2-pf-1633041198065oegccuf7ng").
				Aria("role", "tabpanel").
				TabIndex(0).
				DataSet("ouia-component-type", "PF4/TabContent").
				DataSet("ouia-safe", true).
				Body(
					app.Text("\n    Database\n  "),
				),
			app.Div().
				Style("margin-top", " 20px").
				Body(
					app.Div().
						Class("pf-c-check").
						Body(
							app.Input().
								ID("toggle-box-filled-icon").
								Name("toggle-box-filled-icon").
								Class("pf-c-check__input").
								Type("checkbox").
								Aria("invalid", "false").
								Aria("label", "show box variation checkbox with filled icon tabs").
								DataSet("ouia-component-type", "PF4/Checkbox").
								DataSet("ouia-safe", true).
								DataSet("ouia-component-id", "OUIA-Generated-Checkbox-9"),
							app.Label().
								Class("pf-c-check__label").
								For("toggle-box-filled-icon").
								Body(
									app.Text("isBox"),
								),
						),
				),
		)
}
