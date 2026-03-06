package main

import "testing"

func TestGreetWithName(t *testing.T) {
	result := greet("DevOps")
	expected := "Hello, DevOps!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGreetEmpty(t *testing.T) {
	result := greet("")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}
