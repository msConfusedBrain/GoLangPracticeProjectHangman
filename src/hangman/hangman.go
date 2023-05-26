package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var categories = []string{
	"profession",
	"fruit",
	"vegetable",
	"game",
	"AWS technology",
	"country",
}
var vegetables = []string{"tomato", "onion", "cabbage", "bellpepper", "potato", "cucumber", "broccoli", "ginger", "garlic", "cauliflower", "beetroot", "carrot", }
var fruits = []string{
	"apple",
	"banana",
	"cherry",
	"dates",
	"mango",
	"fig",
	"grape",
	"orange",
	"kiwi",
	"chiku",
}
var professions = []string{"chef", "driver", "engineer", "doctor", "artist", "teacher"}
var countries = []string{"america", "australia", "india", "pakistan", "srilanka", "argentina", "brazil", "newzealand", "ethiopia", "canada", }
var games = []string{"chess", "basketball", "football", "cricket", "tennis", "hockey", "shooting", "boxing", "golf", "volleyball", "baseball"}
var category2wordsmap = make(map[string][]string)

func main() {
	// Initialize the random number generator

	category2wordsmap["profession"] = professions
	category2wordsmap["fruit"] = fruits
	category2wordsmap["vegetable"] = vegetables
	category2wordsmap["game"] = games
	category2wordsmap["country"] = countries
	rand.Seed(time.Now().UnixNano())
	category := categories[rand.Intn(len(categories))]
	fmt.Println("HINT is it a : ", category)
	words := category2wordsmap[category]
	// Choose a random word
	word := words[rand.Intn(len(words))]

	// Set up the game state
	guesses := make(map[string]bool)
	remaining := 6

	// Display the initial game state
	displayState(word, guesses, remaining)

	// Start the game loop
	for remaining > 0 {
		// Ask the player for a guess
		guess := strings.ToLower(getGuess())
		if len(guess)>1{
			fmt.Println("You can only enter one letter at a time")
		}
		if !unicode.IsLetter(rune(guess[0])){
			fmt.Println("You can only enter an alphabet")
		}

		// Check if the guess is correct
		if strings.Contains(word, guess) {
			guesses[guess] = true
		} else {
			remaining--
		}

		// Display the updated game state
		displayState(word, guesses, remaining)

		// Check if the player has won
		if isWin(word, guesses) {
			fmt.Println("Congratulations, you won!")
			return
		}
	}

	// If the player reaches this point, they have lost
	fmt.Println("Sorry, you lose. The word was", word)
}

func displayState(word string, guesses map[string]bool, remaining int) {
	// Display the word with underscores for unguessed letters
	for _, c := range word {
		if guesses[string(c)] {
			fmt.Printf("%c ", c)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()

	// Display the remaining guesses
	fmt.Printf("Remaining guesses: %d\n", remaining)
}

func getGuess() string {
	// Ask the player for a guess
	fmt.Print("Guess a letter: ")
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')
	return strings.TrimSpace(guess)
}

func isWin(word string, guesses map[string]bool) bool {
	// Check if all letters in the word have been guessed
	for _, c := range word {
		if !guesses[string(c)] {
			return false
		}
	}
	return true
}
func uniqueChars(str string) int {
    charMap := make(map[rune]bool)
    for _, char := range str {
        charMap[char] = true
    }
    return len(charMap)
}
