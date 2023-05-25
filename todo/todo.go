package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetTodos() []Todo {
	file, err := os.OpenFile("todos.json", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	size := fi.Size()
	data := make([]byte, size)

	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}

	var todos []Todo

	if err := json.Unmarshal(data, &todos); err != nil {
		todos = []Todo{}

		newData, err := json.Marshal(todos)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("todos.json", newData, 0644)
		if err != nil {
			panic(err)
		}

		return todos
	}

	return todos
}

func AddTodo(todos []Todo, name string, description string) {
	for _, todo := range todos {
		if todo.Name == name {
			fmt.Printf("A todo with the name %s already exists. Please choose a different name.\n", name)

			return
		}
	}

	newTodo := Todo{
		Name:        name,
		Description: description,
	}

	todos = append(todos, newTodo)

	newData, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("todos.json", newData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New todo item added: %s.\n", name)
}

func RemoveTodo(todos []Todo, name string) {
	for i, todo := range todos {
		if todo.Name == name {
			todos = append(todos[:i], todos[i+1:]...)

			newData, err := json.Marshal(todos)
			if err != nil {
				panic(err)
			}

			err = os.WriteFile("todos.json", newData, 0644)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Todo item %s removed successfully.\n", name)

			return
		}
	}

	fmt.Printf("The todo item %s doesn't exist. Please double-check the name and try again.\n", name)
}
