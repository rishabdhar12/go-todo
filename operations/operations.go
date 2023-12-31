package operations

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lithammer/shortuuid"
	"github.com/rishabdhar12/go-todo/models"
	table "github.com/rishabdhar12/go-todo/tables"
)

func CreateNewTodo(list []models.TodoModel) []models.TodoModel {
	id := shortuuid.New()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter todo : ")
	scanner.Scan()
	iTodo := scanner.Text()

	fmt.Printf("Enter Due Date : ")
	scanner.Scan()
	iDueDate := scanner.Text()

	fmt.Printf("Enter description : ")
	scanner.Scan()
	iDescription := scanner.Text()

	validate := ValidateInput(iTodo, iDueDate, iDescription)
	if !validate {
		panic("Validation Error")
	}

	item := models.TodoModel{
		ID:          id,
		Name:        iTodo,
		DueDate:     iDueDate,
		Status:      false,
		Description: iDescription,
	}

	list = append(list, item)

	WriteToFile(list, item)

	return list
}

func ListTodos() {
	file, err := os.ReadFile("todos.json")
	if err == nil {

		var list []models.TodoModel

		err = json.Unmarshal(file, &list)
		if err != nil {
			panic(err)
		}

		table.PrintTable(list)
	} else {
		fmt.Println("No data found")
	}

}

func UpdateTodo() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter the id of the completed todo : ")
	scanner.Scan()
	iID := scanner.Text()

	if iID == "" {
		panic("ID cannot be blank")
	}

	file, err := os.ReadFile("todos.json")

	if err != nil {
		panic("File not found")
	}

	var list []models.TodoModel

	err = json.Unmarshal(file, &list)

	if err != nil {
		errString := fmt.Sprintf("File is empty : %s", err)
		panic(errString)
	}

	for i, item := range list {
		if item.ID == iID {
			list[i].Status = true
			break
		} else {
			errString := fmt.Sprintf("ID doesn't exist : %s", err)
			panic(errString)
		}
	}

	updatedData, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("todos.json", updatedData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Todo updated")
}

func DeleteTodo() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter the id of the todo you want to delete : ")
	scanner.Scan()
	iID := scanner.Text()

	if iID == "" {
		panic("ID cannot be blank")
	}

	file, err := os.ReadFile("todos.json")

	if err != nil {
		panic("File not found")
	}

	var list []models.TodoModel

	err = json.Unmarshal(file, &list)

	if err != nil {
		errString := fmt.Sprintf("File is empty : %s", err)
		panic(errString)
	}

	for i, item := range list {
		if item.ID == iID {
			list = append(list[:i], list[i+1:]...)
			break
		} else {
			errString := fmt.Sprintf("ID doesn't exist : %s", err)
			panic(errString)
		}
	}

	updatedData, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("todos.json", updatedData, 0644)
	if err != nil {
		panic(err)
	}
}

func ValidateInput(iName string, iDueDate string, iDescription string) bool {
	if iName == "" || iDueDate == "" || iDescription == "" {
		return false
	} else {
		return true
	}
}

func WriteToFile(list []models.TodoModel, item models.TodoModel) {
	file, err := os.ReadFile("todos.json")

	if err == nil {
		err = json.Unmarshal(file, &list)
		if err != nil {
			panic(err)
		}

		list = append(list, item)

		updtedData, err := json.Marshal(list)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("todos.json", updtedData, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := os.Create("todos.json")
		if err != nil {
			panic(err)
		}

		updtedData, err := json.Marshal(list)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("todos.json", updtedData, 0644)
		if err != nil {
			panic(err)
		}

	}
}

func PrintHelp() {
	fmt.Println("A simple CLI app to manage tyour todos")
	fmt.Println("Usage: ./todo -[arguments]")
	fmt.Println("Commands:")
	fmt.Println("a\t\t\t\tAdd a new todo")
	fmt.Println("d\t\t\t\tDelete a todo")
	fmt.Println("u\t\t\t\tUpdate a todo")
	fmt.Println("l\t\t\t\tList todos")
	fmt.Println("h\t\t\t\tPrint help")
	fmt.Println("v\t\t\t\tPrint version")
}

func PrintVersion() {
	fmt.Println("Version 1.0.0")
}
