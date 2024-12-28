package main

import (
	"bufio"   // reading input from the user
	"fmt"     // formated I/O to the console
	"os"      // interacts with OS: reads command-line arguments, open files, handles errors
	"strings" // manipulates strings: trimming, splitting
)

// ANSI colour codes
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type Note struct {
	Title     string // The ID of the note
	Body      string // The content of the note
	Timestamp string // The timestamp when the note was saved
}

// Variables to hold the collection file name and its password file name
var collectionName string
var passwordFileName string

// savePassword stores the password for a specific collection
func savePassword(password string) error {
	file, err := os.Create(passwordFileName) // Create password file specific to the collection
	if err != nil {
		return err
	}
	defer file.Close() // close file after operation
	_, err = fmt.Fprintln(file, password) // Write the password to the file
	return err
}

// passwordExists checks if a password file exists for the collection
func passwordExists() bool {
	_, err := os.Stat(passwordFileName) // Check if the password file exists for the collection
	return !os.IsNotExist(err)
}

// verifyPassword compares the entered password with the stored password for the collection
func verifyPassword(password string) bool {
	file, err := os.Open(passwordFileName) // Open the password file for the collection
	if err != nil {
		return false
	}
	defer file.Close()

	var storedPassword string
	fmt.Fscanln(file, &storedPassword) // Read the stored password from the file
	return strings.TrimSpace(storedPassword) == password // Compare the stored password with the entered one
}

func main() {
	// If no argument or wrong number of arguments or help, display help message
	if len(os.Args) != 2 || strings.ToLower(os.Args[1]) == "help" {
		fmt.Println("\nThis tool helps manage a collection of notes.")
		fmt.Println("You can view, add, and delete notes in the specified collection.")
		fmt.Println("If the collection doesn't exist, it will be created automatically.")
		fmt.Println("\n\033[32mUsage: ./notestool <collection_name>\033[0m")
		fmt.Println("\nRead \033[1mREADME.md\033[0m to learn more about the tool.\n")
		return
	}

	collectionName = os.Args[1] + ".txt" // Set the collection file name (e.g., "coding_ideas.txt")
	passwordFileName = os.Args[1] + "_password.txt" // Set the password file name specific to the collection

	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("\n\033[1mWelcome to the Notes Tool!\033[0m")

	// Check if the collection has a password file
	if !passwordExists() {
		// Loop until the user enters a non-empty password
		for {
			fmt.Print("Set a password for this collection: ")
			password, _ := reader.ReadString('\n') // Ask the user to set a password for this collection
			password = strings.TrimSpace(password) // Trim any unnecessary spaces
			
			// Check if the password is empty
			if password == "" {
				fmt.Println("\033[31mPassword cannot be empty. Please enter a valid password.\033[0m")
			} else {
				savePassword(password) // Save the password to the collection's password file
				fmt.Println("Password set. Please restart the program and use this password to access the collection.")
				return
			}
		}
	}

	// If the collection already has a password, prompt the user to enter it
	fmt.Print("Enter password for this collection: ")
	password, _ := reader.ReadString('\n') // Prompt the user for the collection password
	password = strings.TrimSpace(password)

	// Verify the entered password against the stored one
	if !verifyPassword(password) {
		fmt.Println("Incorrect password. Access denied.")
		return
	}

	// Once the correct password is entered, proceed to show the menu
	for {
		// Main menu options
		fmt.Println("\nSelect operation (1-4):")
		fmt.Println("1. Show Notes")
		fmt.Println("2. Add Note")
		fmt.Println("3. Delete Note")
		fmt.Println("4. Exit")

		// Read user's choice
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Trim whitespace
		choice := strings.ToLower(input) // Convert to lowercase

		switch choice {
		case "1":
			showNotes() // Show notes in the collection
		case "2":
			addNote() // Add a new note to the collection
		case "3":
			deleteNote() // Delete a note from the collection
		case "4":
			fmt.Println("\033[1mGoodbye!\033[0m")
			return
		default:
			fmt.Println(Red + "Invalid input. Please select 1, 2, 3, or 4." + Reset)
		}
	}
}
