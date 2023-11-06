package main

import hangman "Hangman/fonction"

func main() {
	hangman.ShowTextFromFile(hangman.SelectLevel())
	hangman.Menu()
}
