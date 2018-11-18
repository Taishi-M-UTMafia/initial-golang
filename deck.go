package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

// function名とbacketの間のdeckは、returnされるもののtypeを表す
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// indexである変数i, jは使わないので_に置き換える
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// (d, deck)がreceiver
// dはdeckの代わりに使うインスタンス的なもの、rubyでいうself
// printメソッドはtypeがdeckのものに限り使える
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// error？
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	// 処理がうまくいったらerrにはnilが格納される
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ", ")
	return deck(s)
}
