package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
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
<<<<<<< HEAD
	fmt.Println(lines[rand])
=======
	randomIndex := rand.Intn(len(lines))
	Word := lines[randomIndex]
	fmt.Println(Word)
>>>>>>> e61d98012236d1f8fd8843de5c573c3bfa2c38bb
}
func ParsingDico() {

}

// on stock le for range dans une variable et quand y'a un \n on met tout ca dans un tableau
