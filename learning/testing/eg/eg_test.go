package eg_test

import (
	. "github.com/atotto/go/learning/testing/eg"
	"testing"
)

func TestFunc(t *testing.T) {
	result := Foo()
	if result != "foo" {
		t.Errorf("result = %v, want %v", result, "foo")
	}
}
