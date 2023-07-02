package learn

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if HelloWorld("Hello") != "Hello World" {
		t.Error("Test failed")
	}
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("Test failed")
	}
}

func TestMultiply(t *testing.T) {
	if multiply(1, 2) != 2 {
		t.Error("Test failed")
	}
}
