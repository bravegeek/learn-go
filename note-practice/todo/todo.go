package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("Your todo has the following text: \n\n%v\n\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := strings.ReplaceAll(todo.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: text,
	}, nil
}
