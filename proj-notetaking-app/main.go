package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dunky.com/proj-notetaking-app/note"
	"dunky.com/proj-notetaking-app/todo"
)

// Main function
func main() {

	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote.Display() // Calling display func from note package

	err = saveData(todo)
	if err != nil {
		return
	}

	err = saveData(userNote)
	if err != nil {
		return
	}

} // End of main function

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content

}

// Interface implementation
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Print("Saving data failed")
		return err
	}

	fmt.Print("Saving data succeeded!")
	return nil
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

// Interface -> Contract
type saver interface {
	Save() error
}
