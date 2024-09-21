package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Error("Expected deck length to be 20 but got ", len(d))
	}
}
