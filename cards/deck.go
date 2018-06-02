package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type deck []string

func createNewDeck() deck {
	var newDeck deck
	suits := []string{"Diamond", "Hearts", "Spades", "Clovers"}
	cards := []string{"Ace", "King", "Queen", "Jack"}

	for _, suit := range suits {
		for _, card := range cards {
			newDeck = append(newDeck, card+" of "+suit)
		}
	}
	return newDeck
}

func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), 0777)
	if err != nil {
		fmt.Println("Something went wrong in saving file.")
		return err
	}
	fmt.Println("Deck saved to file " + fileName)
	return nil
}

func loadDeckFromFile(fileName string) deck {
	deck, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return strings.Split(string(deck), ",")
}
