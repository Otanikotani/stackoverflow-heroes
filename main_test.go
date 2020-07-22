package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSanity(t *testing.T) {
	fmt.Println("Test Sanity")
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}