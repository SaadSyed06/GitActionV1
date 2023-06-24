// my_test.go

package mypackage

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d, expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Multiply(2, 3) = %d, expected %d", result, expected)
	}
}
