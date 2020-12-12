package tpl

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init("../../template")
	for _, i := range GetTplIndex() {
		if _, err := Execute(i, nil); err != nil {
			t.Fatal(err)
		}
	}
}
