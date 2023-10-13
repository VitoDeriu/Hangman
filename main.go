package main

import hangman "Hangman/fonction"

func main() {
	name := "fonction/noms_monstres.txt"
	hangman.ShowTextFromFile(name)
	//hangman.ParsingDico(name)
}
