package cm

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"text/template"
)

// / given a template, we can replace the placeholders with the specified parameter
func TestTemplate(t *testing.T) {
	tmpl, err := template.New("test").Parse("{{.input}}")
	if err != nil {
		panic(err)
	}

	unit := Job{
		parameter: map[string]string{
			"input": "dog",
		},
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, unit.parameter)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "dog", output.String())
}
