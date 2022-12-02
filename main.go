package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lowerAlphabet     = "abcdefghijklmnopqrstuvwxyz"
	upperAlphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers           = "0123456789"
	specialCharacters = "?!@#$%&*"
	charactersList    = lowerAlphabet + upperAlphabet + numbers + specialCharacters
)

var length int

func main() {
	fmt.Println("Enter password length: ")
	fmt.Scan(&length)

	rand.Seed(time.Now().Unix())

	password := generatePassword(length)
	fmt.Println("Password: " + password)
}

func generatePassword(length int) string {
	var password strings.Builder

	randomSpecialChar := rand.Intn(len(specialCharacters))
	password.WriteString(string(specialCharacters[randomSpecialChar]))

	randomNumber := rand.Intn(len(numbers))
	password.WriteString(string(numbers[randomNumber]))

	randomUpperLetter := rand.Intn(len(upperAlphabet))
	password.WriteString(string(upperAlphabet[randomUpperLetter]))

	randomLowerLetter := rand.Intn(len(lowerAlphabet))
	password.WriteString(string(lowerAlphabet[randomLowerLetter]))

	for i := 0; i < length-4; i++ {
		random := rand.Intn(len(charactersList))
		password.WriteString(string(charactersList[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
