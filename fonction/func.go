package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var Word string
var Tableau []string

func ShowTextFromFile(n string) {

	file, err := os.Open(n)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	//lines := make([]string, 0)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	randomIndex := rand.Intn(len(lines))
	Word := lines[randomIndex]
	for _, i := range Word {
		if i == '-' {
			Tableau = append(Tableau, ("-"))
		} else {
			Tableau = append(Tableau, ("_"))
		}
	}
	fmt.Println(Word)
	fmt.Println(Tableau)
	Inputplayer(Tableau)
}

func Inputplayer(n []string) {
	var input string
	fmt.Println("\nEntrez une lettre")
	fmt.Scanln(&input)
	for i := 0; i < len(Tableau); i++ {
		if string(i) == input {
			Tableau = append(Tableau, (input))
		}
	}
	fmt.Println(Tableau)

}

// on stock le for range dans une variable et quand y'a un \n on met tout ca dans un tableau
