package main

import (
	"testing"
)

// Тест на установку бита в 1
func TestSetBit(t *testing.T) {
	var pos int64
	num := int64(0)
	pos = 0
	bit := 1
	expected := int64(1)
	if result := ChangeBit(num, pos, int64(bit)); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Тест на установку бита 0 когда он уже 0
func TestSetBit1(t *testing.T) {
	var pos int64
	num := int64(0)
	pos = 0
	bit := 0
	expected := int64(0)
	if result := ChangeBit(num, pos, int64(bit)); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Тест на установку бита 1 когда он уже 1
func TestSetBit2(t *testing.T) {
	var pos int64
	num := int64(1)
	pos = 0
	bit := 1
	expected := int64(1)
	if result := ChangeBit(num, pos, int64(bit)); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Тест на сброс бита до 0
func TestClearBit(t *testing.T) {
	var pos int64
	num := int64(1)
	pos = 0
	bit := 0
	expected := int64(0)
	if result := ChangeBit(num, pos, int64(bit)); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
