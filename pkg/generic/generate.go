package generic

import (
	"io"
	"text/template"
)

type nameAndPrimitiveKeyword struct {
	Name             string
	PrimitiveKeyword string
}

func Generate(nameMap map[string]string, source *template.Template, writer io.Writer) (err error) {
	templateVars := nameAndPrimitiveKeyword{}
	for name, primitiveKeyword := range nameMap {
		templateVars.Name = name
		templateVars.PrimitiveKeyword = primitiveKeyword
		err = source.Execute(writer, templateVars)
		if err != nil {
			return
		}
	}
	return
}
