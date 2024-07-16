package main

import (
	"bytes"
	"fmt"
	"os"
	"golang.org/x/term"
	"github.com/Simon-Busch/go-file-encrypt/fileEncrypt"
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
	fmt.Println("  encrypt: Encrypt a file /path/to/file")
	fmt.Println("  decrypt: Decrypt a file")
	fmt.Println("  help: Print this help")
}

func encryptHandle() {
	fmt.Println("Encrypting...")
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file to encrypt, missing the path")
		os.Exit(0)
	}

	file := os.Args[2]
	if (!validateFile(file)) {
		fmt.Println("File is valid")
	} else {
		fmt.Println("File is invalid")
	}

	password := getPassword()
	fmt.Println("\n Encrypting file: ", file)
	filecrypt.Encrypt(file, password)
	fmt.Println("File encrypted")
}

func decryptHandle() {
	fmt.Println("Decrypting...")
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file to encrypt, missing the path")
		os.Exit(0)
	}

	file := os.Args[2]
	if (!validateFile(file)) {
		fmt.Println("File is valid")
	} else {
		fmt.Println("File is invalid")
	}

	fmt.Println("Enter password:")
	password, _ := term.ReadPassword(0)
	fmt.Println("\n Decrypting file: ", file)
	filecrypt.Decrypt(file, password)
	fmt.Println("File decrypted successfuly")
}

func getPassword() []byte {
	fmt.Println("Enter password: ")
	password, _ := term.ReadPassword(0)

	fmt.Println("\n Confirm password: ")

	confirmPassword, _ := term.ReadPassword(0)

	if !validatePassword(password, confirmPassword) {
		fmt.Println("Passwords do not match")
		fmt.Println("Please try again")
		getPassword()
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

// Check if file exist
func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true;
}
