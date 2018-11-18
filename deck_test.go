package main

import "testing"

// 引数のt~~は
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		// Errorf, Printfなどで文字列に直接変数などを挿入したい場合、
		// 第一引数の語尾に'%v'をつける
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
