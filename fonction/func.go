package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var Word string

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
	fmt.Println(lines[rand])
}
func ParsingDico() {

}

// on stock le for range dans une variable et quand y'a un \n on met tout ca dans un tableau
