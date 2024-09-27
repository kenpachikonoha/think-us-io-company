package main

import (
	"slices"
	"testing"
)

func TestSumZero(t *testing.T) {
	data := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	expected := []int{5, 8}
	got := sumZero(data)
	if slices.Compare(expected, got) != 0 {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
