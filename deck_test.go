package main

import (
	"os"
	"testing"
)

// 引数のt~~は
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		// Errorf, Printfなどで文字列に直接変数などを挿入したい場合、
		// 第一引数の語尾に'%v'をつける
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %v ", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card is Four of Clubs, but got %v ", d[len(d)-1])
	}
}

// おいおい元の関数を探しやすいようにテストの関数名を長くしてある
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards indeck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
