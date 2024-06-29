package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Volteal/todo-app/internal/models"
	"github.com/Volteal/todo-app/internal/utilities"
)

const (
	todoFile = "./todo-list.json"
)

func main() {
	list := flag.Bool("list", false, "List all todo items!")
	add := flag.Bool("add", false, "Add a new todo item!")
	complete := flag.Int("complete", 0, "Mark a todo item as complete!")
	del := flag.Int("del", 0, "Delete a todo item as complete!")

	flag.Parse()

	todoList := &models.TodoList{}

	if err := todoList.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *list:
		todoList.Show()
	case *add:
		taskName, taskNote, err := utilities.GetUserInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todoList.Add(taskName, taskNote)
		err = todoList.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := todoList.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todoList.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0:
		err := todoList.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todoList.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid command given")
		os.Exit(0)
	}
}
