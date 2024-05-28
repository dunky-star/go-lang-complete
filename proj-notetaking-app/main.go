package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dunky.com/note/note"
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
} // End of main function

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content

}

func getUserInput(promptText string) string {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
