package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func deleteNote() {
	reader := bufio.NewReader(os.Stdin) // reads input
	fmt.Print("\n\033[1mEnter note ID to delete (press 0 to cancel):\033[0m ")
	targetID, _ := reader.ReadString('\n')
	targetID = strings.TrimSpace(targetID)

	// cancel if user pressed 0
	if targetID == "0" {
		fmt.Println("\033[32mDelete operation cancelled.\033[0m")
		return
	}

	file, err := os.Open(collectionName) // open collection file
	if err != nil {
		fmt.Println("Error reading notes:", err)
		return
	}
	defer file.Close()

	var newLines []string // slice to hold all the lines that are not deleted
	found := false // flag that tracks if a note's ID was found and deleted

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // scans file line-by-line and returns the current line as a string
		line := scanner.Text()
		parts := strings.SplitN(line, " - ", 2)
		if len(parts) == 2 {
			id := parts[0]
			if id == targetID { // check for target ID
				fmt.Println("\033[31mDeleted:\033[0m", line)
				found = true // mark as found
				continue // skip adding to newLines slice
			}
		}
		newLines = append(newLines, line) // add remeining notes to slice
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading notes:", err)
		return
	}

	if !found {
		fmt.Println("\033[31mError: No note found with the given ID.\033[0m")
		return
	}

	file, err = os.Create(collectionName)
	if err != nil {
		fmt.Println("Error saving notes:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range newLines { // iterates through newLines and writes them into the file
		if _, err := writer.WriteString(line + "\n"); err != nil {
			fmt.Println("Error writing notes:", err)
			return
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error saving notes:", err)
	} else {
		fmt.Println("\033[32mNotes updated!\033[0m")
	}
}
