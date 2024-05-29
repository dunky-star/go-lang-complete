package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dunky.com/proj-notetaking-app/note"
)

// Main function
func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}
	userNote.Display() // Calling display func from note package
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
} // End of main function

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content

}

func getUserInput(promptText string) string { // Handling long string
	fmt.Printf("%v ", promptText)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
