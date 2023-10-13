package hangman

import (
	"fmt"
	"os"
)

var List []string

func ShowTextFromFile(n string) {

	content, err := os.ReadFile(n)
	if err != nil {
		fmt.Println("Erreur")
	} else {
		List = append(List, string(content))
	}
	fmt.Println(List)
}

func ParsingDico(a string) {

	for _, dicoParse := range a {
		var mot []string
		if dicoParse == '\n' {
			mot = append(mot, string(dicoParse))
			fmt.Println(mot)
		}
	}
	return
}

// on stock le for range dans une variable et quand y'a un \n on met tout ca dans un tableau
