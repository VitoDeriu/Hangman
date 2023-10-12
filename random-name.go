package Hangman

import ("os"
"fmt" )

func ShowTextFromFile(n string){
	var list []string
content, err := os.ReadFile(n)
if err != nil {
	fmt.Println("Erreur")
} else {
	list = append(list, string(content))
}
fmt.Println(list)
}

