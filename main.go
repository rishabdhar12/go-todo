package main

import (
	"flag"

	"github.com/rishabdhar12/go-todo/models"
	"github.com/rishabdhar12/go-todo/operations"
)

func main() {

	var list []models.TodoModel

	HelpFlag := flag.Bool("h", false, "help")
	VersionFlag := flag.Bool("v", false, "Version")
	AddTodoFlag := flag.Bool("a", false, "Add Todo")
	DeleteTodoFlag := flag.Bool("d", false, "Delete Todo")
	UpdateTodoFlag := flag.Bool("u", false, "Update Todo")
	ListTodoFlag := flag.Bool("l", false, "List Todo")

	flag.Parse()

	if *AddTodoFlag {
		operations.CreateNewTodo(list)
	} else if *DeleteTodoFlag {
		operations.DeleteTodo()
	} else if *UpdateTodoFlag {
		operations.UpdateTodo()
	} else if *ListTodoFlag {
		operations.ListTodos()
	} else if *HelpFlag {
		operations.PrintHelp()
	} else if *VersionFlag {
		operations.PrintVersion()
	}
}
