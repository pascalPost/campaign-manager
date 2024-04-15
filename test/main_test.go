package test

import (
	cm "github.com/campaign-manager/internal"
	"testing"
)

func TestUnitCreation(t *testing.T) {
	param := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	unit := cm.NewJob("template.txt", param, "output")
	cm.Generate(*unit)
}
