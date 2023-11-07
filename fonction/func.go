package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var Word string       //mot a deviner
var TabUnder []string //tableau d'underscore
var win bool          //verif si on a win (si y'a plus d'underscore)
var Guessed []string  //liste des lettres déjà rentrer
var Graph int         //compteur de point pour le graph

func ShowTextFromFile(path string) {
	file, err := os.Open(path)
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
	randomIndex := rand.Intn(len(lines))
	Word = ToLower(lines[randomIndex])
	Underscore(Word)
	RandomLetter()
}

func Underscore(Word string) {
	for _, i := range Word {
		if i == '-' {
			TabUnder = append(TabUnder, ("-"))
		} else {
			TabUnder = append(TabUnder, ("_"))
		}
	}
}

func SelectLevel() string {
	var Input string
	fmt.Println("Bienvenu dans Hangman, veuillez choisir votre difficultée:")
	fmt.Println("1 : Facile ")
	fmt.Println("2 : Intermédiaire ")
	fmt.Println("3 : Difficile ")
	fmt.Println("4 : Horror ")
	fmt.Scanln(&Input)
	switch Input {
	case "1":
		return "dico_folder/dico_facile.txt"
	case "2":
		return "dico_folder/dico_intermediaire.txt"
	case "3":
		return "dico_folder/dico_difficile.txt"
	case "4":
		return "dico_folder/noms_monstres.txt"
	}
	return ""
}

func Menu() {
	if win {
		DisplayHangman()
		fmt.Printf("Bien jouer, vous avez deviner le mot : %s", Word)
		return
	} else if Graph > 9 {
		fmt.Printf("Vous avez perdu, le mot à deviner était : %s", Word)
		return
	} else {
		Display()
		WordOrLetter()
		Menu()
	}
}

func Display() {
	fmt.Print("Voici le mot à deviner :")
	for _, i := range TabUnder {
		fmt.Print(i, " ")
	}
	fmt.Println("")
}

func WordOrLetter() {
	var Input string
	fmt.Println("1 : Entrez le mot entier")
	fmt.Println("2 : Entrez une seule lettre")
	fmt.Scan(&Input)
	fmt.Scan()
	switch Input {
	case "1":
		fmt.Println("Entrez un mot :")
		fmt.Scan(&Input)
		if IsInGuessed(Input, Guessed) {
			fmt.Println("Vous avez déjà essayé ce mot")
			return
		}
		IsTheWord(Input)
	case "2":
		fmt.Println("Entrez une lettre :")
		fmt.Scan(&Input)
		if IsInGuessed(Input, Guessed) {
			fmt.Println("Vous avez déjà essayé cette lettre")
			return
		}
		IsInWord(Input)
		return
	default:
		fmt.Println("Choix invalide")
		WordOrLetter()
	}
}

func IsTheWord(w string) {
	if w == Word {
		win = true
		return
	} else {
		Guessed = append(Guessed, w)
	}
	win = false
	Graph += 2
	DisplayHangman()
}

func IsInGuessed(g string, Tab []string) bool {
	for _, i := range Tab {
		if i == g {
			return true
		}
	}
	return false
}

func IsComplete() {
	for _, i := range TabUnder {
		if i == "_" {
			win = false
			return
		}
	}
	win = true
}

func IsInWord(l string) {
	Guessed = append(Guessed, l)
	for _, i := range Word {
		if string(i) == l {
			ChangeTableau(l)
			IsComplete()
			return
		}
	}
	fmt.Println("la lettre n'est pas dans le mot...")
	Graph++
	DisplayHangman()
}

func ChangeTableau(c string) {
	for id, i := range Word {
		if string(i) == c {
			TabUnder[id] = c
		}
	}
}
func RandomLetter() {
	if len(Word) > 5 {
		id := rand.Intn(len(Word))
		ChangeTableau(string(Word[id]))
		Guessed = append(Guessed, string(Word[id]))
	}
	if len(Word) > 7 {
		ind := rand.Intn(len(Word))
		ChangeTableau(string(Word[ind]))
		Guessed = append(Guessed, string(Word[ind]))
	}
}

func ToLower(s string) string {
	var n string
	for _, c := range s {
		if c > 64 && c < 91 {
			n += string(c + 32)
		} else {
			n += string(c)
		}
	}
	return n
}

func DisplayHangman() {
	if Graph == 0 {
		Graph = 1
	}
	file, err := os.Open("GraphHangman/hangman.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer file.Close()
	var lines []string
	for i := 0; i < 7; i++ {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}
	set := Graph * 10
	for i := 10 * (Graph - 1); i < set; i++ {
		fmt.Println(lines[i])
	}
}
