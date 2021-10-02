package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/html2goapp/pkg/components"
)

func main() {
	// Client-side code
	{
		app.Route("/", &components.Home{})
		app.RunWhenOnBrowser()
	}

	// Server-/build-side code

	// Parse the flags
	serve := flag.Bool("serve", false, "Serve the app instead of building it")
	laddr := flag.String("laddr", "0.0.0.0:21255", "Address to listen on when serving the app")
	dist := flag.String("dist", "out/pwa/web", "Directory to build the app to")
	prefix := flag.String("prefix", "/html2goapp", "Prefix to build the app for")

	flag.Parse()

	// Define the handler
	h := &app.Handler{
		Title:           "HTML to go-app Converter",
		Name:            "HTML to go-app Converter",
		ShortName:       "html2goapp",
		Description:     "Convert HTML markup to go-app.dev's syntax.",
		LoadingLabel:    "Convert HTML markup to go-app.dev's syntax.",
		Author:          "Felicitas Pojtinger",
		ThemeColor:      "#151515",
		BackgroundColor: "#151515",
		Icon: app.Icon{
			Default: "/web/default.png",
			Large:   "/web/large.png",
		},
		Keywords: []string{
			"html-converter",
			"code-generation",
			"go-app",
		},
		RawHeaders: []string{
			`<meta property="og:url" content="https://pojntfx.github.io/html2goapp/">`,
			`<meta property="og:title" content="HTML to go-app Converter">`,
			`<meta property="og:description" content="Convert HTML markup to go-app.dev's syntax.">`,
			`<meta property="og:image" content="https://pojntfx.github.io/html2goapp/web/default.png">`,
		},
		Styles: []string{
			"https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly.css",
			"https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly-addons.css",
		},
	}

	// Serve if specified
	if *serve {
		log.Println("Listening on", *laddr)

		if err := http.ListenAndServe(*laddr, h); err != nil {
			log.Fatal("could not serve:", err)
		}

		return
	}

	// Build if not specified
	h.Resources = app.GitHubPages(*prefix)

	if err := app.GenerateStaticWebsite(*dist, h); err != nil {
		log.Fatal("could not build static website:", err)
	}
}
