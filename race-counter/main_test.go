package main

import (
	"testing"
)

func TestRace(t *testing.T) {
	expected := 100000
	got := routine()
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
