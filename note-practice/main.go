package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver // embeds existing interface in another interface
	Display()
}

// type displayer interface {
// 	Display()
// }

// type outputable interface {
// 	Save() error
// 	Display()
// }

func main() {

	// get data
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	// create, display, save Todo
	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	// userTodo.Display()
	// err = saveData(userTodo)

	outputData(userTodo)

	if err != nil {
		return
	}

	// create, display, save Note
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}

	// err = userNote.Save()
	// if err != nil {
	// 	fmt.Println("Saving the note failed.")
	// 	return
	// }

	// fmt.Println("Saving the note succeeded")
}

func outputData(data outputable) {
	data.Display()
	saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed")
		return err
	}

	fmt.Println("Saving the data succeeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
