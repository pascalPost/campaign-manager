package task_creation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInput(t *testing.T) {
	variantsYaml := `
templates:
  - template1
  - template2
matrix:
  key1:
    - v1.1
    - v1.2
  key2:
    - v2.1
    - v2.2
exclude:
  -
    key1: v1.1
    key2: v2.1
include:
  -
    key2: v2.3
`

	config, err := ParseInput([]byte(variantsYaml))
	assert.NoError(t, err)
	assert.Equal(t, &Input{
		Templates: []string{"template1", "template2"},
		Matrix: map[string][]string{
			"key1": {"v1.1", "v1.2"},
			"key2": {"v2.1", "v2.2"},
		},
		Exclude: []map[string]string{
			{"key1": "v1.1", "key2": "v2.1"},
		},
		Include: []map[string]string{
			{"key2": "v2.3"},
		},
	}, config)
}
