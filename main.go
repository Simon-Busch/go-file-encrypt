package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
		case "help":
			printHelp()
		case "encrypt":
			encryptHandle()
		case "decrypt":
			decryptHandle()
		default:
			fmt.Println("Run encrypt or decrypt")
			os.Exit(1)
	}
}


func printHelp() {
	fmt.Println("Usage: ")
	fmt.Println("  go run main.go <command> [args]")
	fmt.Println("Commands: ")
	fmt.Println("  encrypt: Encrypt a file")
	fmt.Println("  decrypt: Decrypt a file")
	fmt.Println("  help: Print this help")
}

func encryptHandle() {
	fmt.Println("Encrypting...")
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file to encrypt, missing the path")
		os.Exit(1)
	}
}

func decryptHandle() {
	fmt.Println("Decrypting...")
}

func getPassword() {
	fmt.Println("Enter password: ")
}

func validatePassword() {

}

func validateFile() {

}
