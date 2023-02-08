package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"

	"github.com/atotto/clipboard"
)

const Help string = `
usage: 
./pwdgen -flag

list of flags:

-u: usage
-h: list of flags
-v: version
`
const Usage string = `
Usage:
./pwdgen word number

example:
./pwdgen hello 123
`

const Version string = `
pwdgen current version: 0.2v

SPDX-License-Identifier: MIT  
Copyright (c) 2023 [Joseph Anthony Debono](joe@jadebono.com)

Github: 

`
const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{};':\"\\|,.<>/?`~"

func main() {

// check os.Args for number of arguments supplied
	if len(os.Args) == 1 {
		fmt.Print("pwdgen has been run either without the arguments or without the flag! Program will terminate here!\nFor help, run pwdgen with one of these flags:\n", Help)
	// if the array is not empty, check for flags from doc.go package
	} else if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-u":
			fmt.Println(Usage)
		case "-v":
			fmt.Println(Version)
		case "-h":
			fmt.Println(Help)
		// if the array is not empty but there are no flags, validate the input
		default:
			fmt.Print("pwdgen has been run without either the word or the number! Program will terminate here!\nFor help, run pwdgen with one of these flags:\n", Help)
		} 
		} else {
		// Get the phrase and number from the command line arguments
		phrase := os.Args[1]
		number := os.Args[2]

		// Concatenate the phrase and number
		input := phrase + number

		// Generate the SHA-256 hash of the concatenated string
		hash := sha256.Sum256([]byte(input))

		// Encode the hash as a hexadecimal string
		hashString := hex.EncodeToString(hash[:])

		// Use the first 16 characters of the hash string as a seed for the random number generator
		seed, _ := binary.Varint([]byte(hashString[:16]))
		rand.Seed(seed)

		// Generate a 128-character string of mixed capital letters, lowercase letters, numbers, and symbols
		result := ""
		for i := 0; i < 128; i++ {
			result += string(characters[rand.Intn(len(characters))])
		}

		// Output the result to the terminal
		fmt.Println(result)

		// Write the result to the clipboard
		err := clipboard.WriteAll(result)
		if err != nil {
			fmt.Println("Error copying to clipboard:", err)
			}

	}
}
