package hangman

import (
	"fmt"
	"os"
)

var List string

func ShowTextFromFile(n string) {

	content, err := os.ReadFile(n)
	if err != nil {
		fmt.Println("Erreur")
	} else {
		List = List + string(content)
	}
	fmt.Println(List)
}

func ParsingDico(a string) {

	for _, dicoParse := range a {

		var tab []string
		var mot string

		mot += string(dicoParse)

		if dicoParse == '\n' {
			tab = append(tab, mot)
			fmt.Println(tab)
		}
	}
	return
}

// on stock le for range dans une variable et quand y'a un \n on met tout ca dans un tableau
