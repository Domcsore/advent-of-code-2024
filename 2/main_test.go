package main

import (
	"testing"
)

func TestIsSafeWithRemovedLevelSafe(t *testing.T) {
	input := []int{7, 6, 4, 2, 1}

	want := true

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestIsSafeWithRemovedLevelSafeTwo(t *testing.T) {
	input := []int{1, 3, 6, 7, 9}

	want := true

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestIsSafeWithRemovedLevelUnsafe(t *testing.T) {
	input := []int{1, 2, 7, 8, 9}

	want := false

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestIsSafeWithRemovedLevelUnsafeTwo(t *testing.T) {
	input := []int{9, 7, 6, 2, 1}

	want := false

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestIsSafeWithRemovedLevelSafeWithRemoval(t *testing.T) {
	input := []int{1, 3, 2, 4, 5}

	want := true

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestIsSafeWithRemovedLevelSafeWithRemovalTwo(t *testing.T) {
	input := []int{8, 6, 4, 4, 1}

	want := true

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestIsSafeWithRemovedLevelSafeWithEndRemoval(t *testing.T) {
	input := []int{8, 6, 4, 2, 3}

	want := true

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestIsSafeWithRemovedLevelSafeWithDirectionFlip(t *testing.T) {
	input := []int{8, 6, 4, 8, 2}

	want := true

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestIsSafeWithRemovedLevelSafeWithFirstChange(t *testing.T) {
	input := []int{900, 6, 4, 3, 2}

	want := true

	got := IsSafeWithRemovedLevel(input)

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}
