package test

import (
	"github.com/campaign-manager/src"
	"testing"
)

func TestUnitCreation(t *testing.T) {
	param := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	unit := cm.NewUnit("template.txt", param, "output")
	cm.Generate(*unit)
}
