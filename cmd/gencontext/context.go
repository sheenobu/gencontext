package main

import (
	"os"
	"strings"

	"text/template"
)

type contextData struct {
	Invocation        string
	AdditionalImports string

	Package string
	Type    string
}

func (g *Generator) generateContext(objectType string) {

	data := contextData{
		Invocation: strings.Join(os.Args[1:], " "),
		Package:    g.pkg.name,
		Type:       objectType,
	}

	t := template.Must(template.New("context").Parse(contextTemplate))

	err := t.Execute(&g.buf, data)
	if err != nil {
		panic(err)
	}
}
