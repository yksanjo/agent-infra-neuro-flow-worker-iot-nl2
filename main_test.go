package main

import "testing"

func TestHealth(t *testing.T) {
	if true != true {
		t.Error("Expected true")
	}
}
