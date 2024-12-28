package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings" // converts strings to int for handling notes' IDs
	"time"	// provides functions for working with timestamps
)

// Function to generate the next available note ID
func getNextID() (string, error) {
	file, err := os.Open(collectionName) // open file of collection
	if err != nil {
		if os.IsNotExist(err) {
			return "001", nil
		}
		return "", err
	}
	defer file.Close()

	maxID := 0
	scanner := bufio.NewScanner(file) // read file line-by-line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 3 {
			continue
		}
		id, err := strconv.Atoi(line[:3]) // convert string to integer
		if err == nil && id > maxID {
			maxID = id
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	nextID := fmt.Sprintf("%03d", maxID+1) // add 0 in the beginning
	return nextID, nil
}

// Function to add a new note
func addNote() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n\033[1mEnter note:\033[0m ")
	body, _ := reader.ReadString('\n')
	body = strings.TrimSpace(body) // necessary for correct representation in Show Notes

	if body == "" {
		fmt.Println("\033[31mYou entered an empty note.\033[0m")
		return
	}

	id, err := getNextID()
	if err != nil {
		fmt.Println("Error generating note ID:", err)
		return
	}

	// Generate the timestamp when the note is created
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	note := Note{
		Title:     id,
		Body:      body,
		Timestamp: timestamp,  // Ensure the timestamp is part of the note
	}

	err = saveNote(note)
	if err != nil {
		fmt.Println("Error saving note:", err)
	} else {
		fmt.Println("\033[1mNote saved with ID:\033[0m", id)
	}
}
