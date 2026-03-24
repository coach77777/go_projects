package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error getting note data:", err)
		return
	}

// Call the Display method to show the note
	userNote.Display()
	err = userNote.Save()

	if err !=nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {

	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {

	fmt.Printf("%v ",prompt)

	// for one word or numbers -small
	// var value string
	// fmt.Scan(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
