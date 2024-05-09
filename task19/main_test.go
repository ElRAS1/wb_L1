package main

import (
	"testing"
)

func TestUtf(t *testing.T) {
	first := "главрыба"
	expected := "абырвалг"
	result, err := revers(first)

	if err != nil {
		t.Fatal("Revers returning error...")
	}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestString(t *testing.T) {
	first := "hello"
	expected := "olleh"
	result, err := revers(first)

	if err != nil {
		t.Fatal("Revers returning error...")
	}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestError(t *testing.T) {
	first := 6
	_, err := revers(first)

	if err == nil {
		t.Fatal("Revers returning error...")
	}

}
