package task_creation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateFromSimpleMatrix(t *testing.T) {
	config := &Input{
		Templates: []string{"template1", "template2"},
		Matrix: map[string][]string{
			"key1": {"v1.1", "v1.2"},
			"key2": {"v2.1", "v2.2"},
		},
	}

	variants := GenerateVariants(config)
	assert.Equal(t, []map[string]string{
		{"key1": "v1.1", "key2": "v2.1"},
		{"key1": "v1.2", "key2": "v2.1"},
		{"key1": "v1.1", "key2": "v2.2"},
		{"key1": "v1.2", "key2": "v2.2"},
	}, variants)
}

func TestGenerateSimpleExclude(t *testing.T) {
	config := &Input{
		Templates: []string{"template1", "template2"},
		Matrix: map[string][]string{
			"key1": {"v1.1", "v1.2"},
			"key2": {"v2.1", "v2.2"},
		},
		Exclude: []map[string]string{
			{"key1": "v1.1", "key2": "v2.1"},
		},
	}

	variants := GenerateVariants(config)
	assert.Equal(t, []map[string]string{
		{"key1": "v1.1", "key2": "v2.2"},
		{"key1": "v1.2", "key2": "v2.1"},
		{"key1": "v1.2", "key2": "v2.2"},
	}, variants)
}

func TestGenerateComplexExclude(t *testing.T) {
	config := &Input{
		Templates: []string{"template1", "template2"},
		Matrix: map[string][]string{
			"key1": {"v1.1", "v1.2", "v1.3"},
			"key2": {"v2.1", "v2.2"},
			"key3": {"v3.1", "v3.2"},
		},
		Exclude: []map[string]string{
			{"key1": "v1.1", "key3": "v3.1"},
		},
	}

	variants := GenerateVariants(config)
	assert.Equal(t, []map[string]string{
		//{"key1": "v1.2", "key2": "v2.1", "key3": "v3.1"},
		//{"key1": "v1.2", "key2": "v2.2", "key3": "v3.1"},
		//{"key1": "v1.3", "key2": "v2.1", "key3": "v3.1"},
		//{"key1": "v1.3", "key2": "v2.2", "key3": "v3.1"},
		{"key1": "v1.2", "key2": "v2.1", "key3": "v3.2"},
		{"key1": "v1.2", "key2": "v2.2", "key3": "v3.2"},
		{"key1": "v1.3", "key2": "v2.1", "key3": "v3.2"},
		{"key1": "v1.3", "key2": "v2.2", "key3": "v3.2"},
	}, variants)
}
