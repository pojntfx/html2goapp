package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pojntfx/html-to-go-app-converter/pkg/converter"
)

func main() {
	src := flag.String("src", "index.html", "HTML source file to convert")
	goAppPkg := flag.String("goAppPkg", "github.com/maxence-charriere/go-app/v9/pkg/app", "Package to use for go-app")
	pkg := flag.String("pkg", "components", "Package to generate component in")
	component := flag.String("component", "MyComponent", "Name of the component to generate")

	flag.Parse()

	htmlSrc, err := ioutil.ReadFile(*src)
	if err != nil {
		log.Fatal("could not open HTML source file:", err)
	}

	source, err := converter.ConvertHTMLToComponent(
		string(htmlSrc),
		*goAppPkg,
		*pkg,
		*component,
	)
	if err != nil {
		fmt.Print(source)

		log.Fatal("could not convert HTML to component:", err)
	}

	fmt.Print(source)
}
