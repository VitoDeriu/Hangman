package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ShowTextFromFile(n string) {

	file, err := os.Open(n)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
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
