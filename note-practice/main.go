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

func main() {

	printSomething(23)
	printSomething("I printed something")

	// get data
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	// create, display, save Todo
	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)

	if err != nil {
		return
	}

	printSomething(userTodo) // won't do anything

	// create, display, save Note
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", typedVal)
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// 	// default:
	// }
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
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
