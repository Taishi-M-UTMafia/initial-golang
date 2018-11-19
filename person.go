package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// こっちの*はtypeがpersonに対するpointerであることを表す、下の*とは完全に別物
func (pointerToPerson *person) updateName(newFirstName string) {
	// こっちの*はpointerのvalueであることを表す
	(*pointerToPerson).firstName = newFirstName
}
