package main

import (
	"os"
	"testing"
)

func TestCreateNewDeck(t *testing.T) {

	newDeck := createNewDeck()

	if len(newDeck) != 16 {
		t.Errorf("Expected lenght is 16. But got %v", len(newDeck))
	}

	if newDeck[0] != "Ace of Diamond" {
		t.Errorf("Expected first card: Ace of Diamonds. But got %s", newDeck[0])
	}

	if newDeck[len(newDeck)-1] != "Jack of Clovers" {
		t.Errorf("Expected last card: Jack of Clovers. But got %s", newDeck[len(newDeck)-1])
	}
}

func TestSaveToFileAndLoadDeckFromFile(t *testing.T) {

	os.Remove("_deckTesting")

	deck := createNewDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := loadDeckFromFile("_deckTesting")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected length %v. But got %v", len(deck), len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
