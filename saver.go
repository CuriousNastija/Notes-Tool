package main

import (
	"bufio"
	"os"
)

// saveNote appends a new note to the notes collection file.
// The note is saved in the format: "ID - note body [timestamp]"
// Takes struct as an argument and returns error, if no error -> returns nil
func saveNote(note Note) error {
	// Open the file in append mode. Create it if it doesn't exist. Write-only mode
	file, err := os.OpenFile(collectionName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 
	if err != nil {
		return err // Return an error if the file cannot be opened
	}
	defer file.Close() // Ensure the file is closed after the operation, starts function Close after current function finishes

	// Create a buffered writer for efficient writing. Temporary holds data in memory
	writer := bufio.NewWriter(file)

	// Write the note to the file using buffer writer in the format: "ID - note body [timestamp]"
	_, err = writer.WriteString(note.Title + " - " + note.Body + " [" + note.Timestamp + "]\n")
	if err != nil {
		return err // Return an error if the write operation fails
	}

	// Flush the buffered writer to ensure all data is written to the file
	return writer.Flush() // Ensure all buffered data is written to the file
}

