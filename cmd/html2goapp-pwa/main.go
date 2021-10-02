package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/html2goapp/pkg/components"
)

func main() {
	serve := flag.Bool("serve", false, "Serve the app instead of building it")
	laddr := flag.String("laddr", "0.0.0.0:1234", "Address to listen on when serving the app")
	dist := flag.String("dist", "out/pwa/web", "Directory to build the app to")
	prefix := flag.String("prefix", "/html2goapp", "Prefix to build the app for")

	flag.Parse()

	app.Route("/", &components.Home{})
	app.RunWhenOnBrowser()

	h := &app.Handler{
		Name:         "HTML to go-app Converter",
		Description:  "Convert HTML markup to go-app.dev's syntax",
		Author:       "Felix Pojtinger",
		LoadingLabel: "Loading HTML to go-app Converter",
		Icon: app.Icon{
			Default: "/web/default.png",
			Large:   "/web/large.png",
		},
		Styles: []string{
			"https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly.css",
			"https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly-addons.css",
		},
	}

	if *serve {
		log.Println("Serving on", *laddr)

		if err := http.ListenAndServe(*laddr, h); err != nil {
			log.Fatal("could not serve:", err)
		}
	} else {
		h.Resources = app.GitHubPages(*prefix)

		if err := app.GenerateStaticWebsite(*dist, h); err != nil {
			log.Fatal("could not build static website:", err)
		}
	}
}
