package main

import hangman "Hangman/fonction"

func main() {
	name := "noms_monstres.txt"
	hangman.ShowTextFromFile(name)
	hangman.ParsingDico(hangman.List)
}
