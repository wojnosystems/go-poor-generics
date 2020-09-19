package generic

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"text/template"
)

func TestGenerate(t *testing.T) {
	commonTemplate, err := template.New("TestGenerate").Parse(`type My{{.Name}} struct {
    value {{.PrimitiveKeyword}}
}
`)
	if err != nil {
		t.Fatalf("unable to parse the template")
	}
	packageName := "my_package"
	cases := map[string]struct {
		namePrimitives map[string]string
		expected       string
	}{
		"blank": {
			namePrimitives: map[string]string{},
			expected:       "package my_package\n\n",
		},
		"one": {
			namePrimitives: map[string]string{
				"Int": "int",
			},
			expected: `package my_package

type MyInt struct {
    value int
}
`,
		},
		"two": {
			namePrimitives: map[string]string{
				"Int":   "int",
				"Int64": "int64",
			},
			expected: `package my_package

type MyInt struct {
    value int
}
type MyInt64 struct {
    value int64
}
`,
		},
	}

	for caseName, c := range cases {
		var err error
		actual := bytes.NewBuffer([]byte{})
		err = Generate(packageName, c.namePrimitives, commonTemplate, actual)
		if err != nil {
			t.Fatalf("unable to render the template for case with name: \"%s\"", caseName)
		}
		assert.Equal(t, c.expected, actual.String(), caseName)
	}
}
