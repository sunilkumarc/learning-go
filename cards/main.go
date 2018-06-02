package main

func main() {
	myDeck := createNewDeck()
	myDeck.saveToFile("my_deck.txt")

	loadedDeck := loadDeckFromFile("my_deck.txt")
	loadedDeck.printDeck()
}
