package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "A Spades" {
		t.Errorf("Expected first card of A Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "K Clubs" {
		t.Errorf("Expected last card of K Clubs, but got %v", d[len(d)-1])
	}
}
