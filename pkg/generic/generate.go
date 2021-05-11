package generic

import (
	"io"
	"sort"
	"text/template"
)

type nameAndPrimitiveKeyword struct {
	Name             string
	PrimitiveKeyword string
}

func Generate(nameMap map[string]string, source *template.Template, writer io.Writer) (err error) {
	templateVars := nameAndPrimitiveKeyword{}

	sortedKeys := make([]string, 0, len(nameMap))
	for key := range nameMap {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	for _, name := range sortedKeys {
		primitiveKeyword := nameMap[name]
		templateVars.Name = name
		templateVars.PrimitiveKeyword = primitiveKeyword
		err = source.Execute(writer, templateVars)
		if err != nil {
			return
		}
	}
	return
}
