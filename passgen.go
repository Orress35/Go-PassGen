package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	special = "!@#$%"
	charset = ""
	reader  = bufio.NewReader(os.Stdin)
)

func main() {
	if getBool("letters") {
		charset += letters
	}

	if getBool("digits") {
		charset += digits
	}

	if getBool("special") {
		charset += special
	}

	amount := getInt("amount")
	length := getInt("length")

	for i := 0; i < amount; i++ {
		fmt.Println(generatePassword(length))
	}
}

func randChar() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(charset))
	return string(charset[random])
}

func getBool(msg string) bool {
	fmt.Print(msg + " (y/n): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return strings.Contains(text, "y")
}

func getInt(msg string) int {
	fmt.Print(msg + ": ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	i, err := strconv.Atoi(text)
	if err != nil {
		return i
	}
	return -1
}

func generatePassword(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += randChar()
	}
	return result
}
