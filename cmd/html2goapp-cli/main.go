package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pojntfx/html2goapp/pkg/converter"
)

func main() {
	// Parse the flags
	src := flag.String("src", "index.html", "HTML source file to convert")
	goAppPkg := flag.String("goAppPkg", "github.com/maxence-charriere/go-app/v9/pkg/app", "Package to use for go-app")
	pkg := flag.String("pkg", "components", "Package to generate component in")
	component := flag.String("component", "MyComponent", "Name of the component to generate")

	flag.Parse()

	// Open the input file
	htmlInput, err := ioutil.ReadFile(*src)
	if err != nil {
		log.Fatal("could not open HTML source file:", err)
	}

	// Convert to Go
	goOutput, err := converter.ConvertHTMLToComponent(
		string(htmlInput),
		*goAppPkg,
		*pkg,
		*component,
	)
	if err != nil {
		fmt.Print(goOutput)

		log.Fatal("could not convert HTML to component:", err)
	}

	// Output the generated Go source
	fmt.Print(goOutput)
}
