package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to display all notes
func showNotes() {
	file, err := os.Open(collectionName) // Open the collection file
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, print that no notes are found
			fmt.Println("\033[31mNo notes found in this collection.\033[0m")
			return
		}
		// Print error if there is an issue opening the file
		fmt.Println(Red+"Error reading notes collection:"+Reset, err)
		return
	}
	defer file.Close() // Ensure the file is closed after reading

	scanner := bufio.NewScanner(file) // Create a scanner to read the file line by line

	hasNotes := false // Flag to check if there are any notes in the collection

	fmt.Println("\n\033[1mNotes:\033[0m")

	// Loop through the file to display notes
	for scanner.Scan() {
		hasNotes = true // Set flag to true as soon as a note is found
		line := scanner.Text()

		// Split the string into  parts by separator - ID - first part, body, and timestamp (assuming "ID - body [timestamp]" format)
		parts := strings.SplitN(line, " - ", 2)
		if len(parts) == 2 {
			id := parts[0] // ID is our first part
			bodyParts := strings.SplitN(parts[1], " [", 2) // split second part with separator [ to take timestamp
			if len(bodyParts) == 2 {
				body := bodyParts[0]
				timestamp := strings.TrimSuffix(bodyParts[1], "]") // delete trailing ]
				// Print each note's ID, content, and timestamp
				fmt.Printf("\033[1mID:\033[0m %s | \033[1mNote:\033[0m %s | \033[1mSaved on:\033[0m %s\n", id, body, timestamp)
			} else {
				// If the format is incorrect, just print the line
				fmt.Println(line)
			}
		}
	}

	// If no notes were found, display a message indicating an empty collection
	if !hasNotes {
		fmt.Println("\033[31mNo notes found in this collection.\033[0m")
	}

	// Handle any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println(Red+"Error reading notes:"+Red, err)
	}
}
