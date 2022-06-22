package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/lerichardson/wordle/errors"
	"github.com/lerichardson/wordle/src"
)

var (
	Black   string = "\033[1;30m"
	Red     string = "\033[1;31m"
	Green   string = "\033[1;32m"
	Yellow  string = "\033[1;33m"
	Purple  string = "\033[1;34m"
	Magenta string = "\033[1;35m"
	Teal    string = "\033[0;97m"
	White   string = "\033[1;37m"
)

func main() {
	var guess string

	var word string = RandomWord()
	var words []string = append(src.ListAnswers(), strings.ToUpper(word))
	fmt.Println("Welcome to wordle")
	guessData := []map[string][5]string{}
	for guesses := 1; guesses < 7; guesses++ {
		fmt.Printf("\nEnter guess %v\n", guesses)
		_, err := fmt.Scanf("%s", &guess)
		errors.Handle(err)
		if guess == word {
			fmt.Printf("\n%sCongratulations! You guessed correctly%s\n", Green, Black)
			break
		} else {
			for _, value := range words {
				if value == guess {
					vector := ColorVector("Grey")

					// stores whether an index is allowed to cause another index to be yellow
					yellow_lock := [5]bool{}

					for j, guess_letter := range guess {
						for k, letter := range word {
							if guess_letter == letter && j == k {
								vector[j] = "Green"
								// now the kth index can no longer cause another index to be yellow
								yellow_lock[k] = true
								break

							}
						}
					}
					for j, guess_letter := range guess {
						for k, letter := range word {
							if guess_letter == letter && vector[j] != "Green" && yellow_lock[k] == false {
								vector[j] = "Yellow"
								yellow_lock[k] = true
							}
						}
					}
					guessData = append(guessData, map[string][5]string{guess: vector})
					DisplayWord(guess, vector)
				}
			}
		}
	}
}
func DisplayWord(word string, vector [5]string) {
	for i, c := range word {
		switch vector[i] {
		case "Green":
			fmt.Print(Green)
		case "Yellow":
			fmt.Print(Yellow)
		case "Grey":
			fmt.Print(Black)
		}
		fmt.Printf("[%c]", c)
		fmt.Print("\033[m\033[m")
	}
}
func ColorVector(color string) [5]string {
	vector := [5]string{}
	for i := range vector {
		vector[i] = color
	}
	return vector
}
func RandomWord() (word string) {
	var words []string = src.ListAnswers()
	rand.Seed(time.Now().UnixNano())
	var zelength int = len(words)
	var indexnum int = rand.Intn(zelength - 1)
	word = words[indexnum]
	return word
}
