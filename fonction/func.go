package hangman

import (
	"fmt"
	"os"
	"bufio"
	"log"
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
		var mot []string
		if dicoParse == '\n' {
			mot = append(mot, string(dicoParse))
			fmt.Println(mot)
		}
	}
	return
}

// on stock le for range dans une variable et quand y'a un \n on met tout ca dans un tableau
