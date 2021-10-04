package converter

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/dave/jennifer/jen"
	"golang.org/x/net/html"
	"mvdan.cc/gofumpt/format"
)

func convertHTMLToStatements(doc *html.Node, goAppPkg string) (*Statement, error) {
	var crawler func(node *html.Node, nthChild int) (*Statement, error)
	crawler = func(node *html.Node, nthChild int) (*Statement, error) {
		el := Null()

		if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
			// Handle text node
			el = Qual(goAppPkg, "Text").Call(Lit(node.Data))
		} else if node.Type == html.ElementNode && node.DataAtom.String() != "" {
			// Handle complex node
			el = Qual(goAppPkg, formatTag(node.DataAtom.String())).Call()
			if nthChild >= 2 {
				el = Line().Qual(goAppPkg, formatTag(node.DataAtom.String())).Call()
			}

			for _, attr := range node.Attr {
				// Attributes to ignore
				if attr.Key == "gutter" || attr.Key == "onload" {
					continue
				}

				// Handle empty attributes
				var val interface{}
				val = attr.Val
				if val == "" {
					val = true
				}

				// Handle `aria-*` and `data-*` attributes
				key := formatKey(attr.Key)
				if strings.Contains(key, "-") {
					parts := strings.Split(key, "-")

					key = formatKey(formatTag(parts[0]))

					el.Dot(key)

					if val == "true" {
						el.Call(Lit(strings.Join(parts[1:], "-")), Lit(true))
					} else {
						el.Call(Lit(strings.Join(parts[1:], "-")), Lit(val))
					}

					val = nil
				} else {
					key = formatKey(formatTag(attr.Key))

					el.Dot(key)
				}

				if key == "TabIndex" {
					// Parse ints for `TabIndex`
					v, err := strconv.Atoi(fmt.Sprintf("%v", val))
					if err != nil {
						return nil, err
					}

					el.Call(Lit(v))
				} else if key == "Style" {
					// Convert string representation of CSS in the style tag to multiple calls

					styleParts := strings.Split(fmt.Sprintf("%v", val), ":")

					// style="" or invalid CSS
					if val == true || len(styleParts) <= 1 || len(styleParts)%2 != 0 {
						el.Call(Lit(""), Lit(""))
					} else {
						for i, key := range styleParts {
							if i%2 == 0 {
								if i == 0 {
									el.Call(Lit(key), Lit(styleParts[i+1]))
								} else {
									el.Dot("Style").Call(Lit(key), Lit(styleParts[i+1]))
								}
							}
						}
					}
				} else if key == "AutoComplete" {
					// Parse booleans for `AutoComplete`
					if val == "off" {
						el.Call(Lit(false))
					} else {
						el.Call(Lit(true))
					}
				} else if key == "Spellcheck" {
					// Parse booleans for `Spellcheck`
					if val == "true" {
						el.Call(Lit(true))
					} else {
						el.Call(Lit(false))
					}
				} else if key == "CrossOrigin" {
					// Convert boolean to strings for `CrossOrigin`
					if val == true {
						el.Call(Lit("true"))
					} else {
						el.Call(Lit("false"))
					}
				} else if key == "Class" {
					// Handle empty `Class` attributes
					if val == true {
						el.Call(Lit(""))
					} else {
						el.Call(Lit(val))
					}
				} else if key == "Width" || key == "Height" || key == "Rows" || key == "Cols" {
					// Parse ints for `Width` and `Height`
					v, err := strconv.Atoi(strings.Trim(fmt.Sprintf("%v", val), "px"))
					if err != nil {
						return nil, err
					}

					el.Call(Lit(v))
				} else if val != nil {
					el.Call(Lit(val))
				}
			}
		}

		children := []Code{}
		i := 0
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			// Tags to ignore
			if child.DataAtom.String() != "svg" {
				child, err := crawler(child, i)
				if err != nil {
					return nil, err
				}

				children = append(children, child)

				i++
			}
		}

		if len(children) > 0 {
			el.Dot("").Line().Id("Body").Call(Line().List(append(children, Line())...))
		}

		return el, nil
	}

	return crawler(doc, 0)
}

func formatTag(tag string) string {
	return strings.Join(strings.Fields(strings.Title(strings.ReplaceAll(tag, "-", " "))), "")
}

func formatKey(key string) string {
	if key == "Id" {
		return "ID"
	}

	if key == "Tabindex" {
		return "TabIndex"
	}

	if key == "role" {
		return "aria-role"
	}

	if key == "Data" {
		return "DataSet"
	}

	if key == "Autocomplete" {
		return "AutoComplete"
	}

	if key == "Crossorigin" {
		return "CrossOrigin"
	}

	return key
}

// ConvertHTMLToComponent converts HTML markup to go-app's syntax
func ConvertHTMLToComponent(
	htmlInput,
	goAppPkg,
	componentPkg,
	componentName string,
) (string, error) {
	// Parse HTML input
	root, err := html.Parse(strings.NewReader(htmlInput))
	if err != nil {
		return "", err
	}

	statements, err := convertHTMLToStatements(root.FirstChild.LastChild.FirstChild, goAppPkg)
	if err != nil {
		return "", err
	}

	// Create package
	src := NewFile(componentPkg)

	// Component Struct
	src.Type().Id(componentName).
		Struct(
			Qual(goAppPkg, "Compo"),
		)

	// Render function
	src.Func().
		Params(
			Id("c").Id("*" + componentName),
		).
		Id("Render").
		Params().
		Params(Qual(goAppPkg, "UI")).
		// Generated statements
		Block(Return(statements))

		// Format source code
	if err := os.Setenv("GOFUMPT_SPLIT_LONG_LINES", "on"); err != nil {
		return "", err
	}

	out, err := format.Source([]byte(fmt.Sprintf("%#v", src)), format.Options{})
	if err != nil {
		return fmt.Sprintf("%#v", src), err
	}

	return string(out), nil
}
