package models

type TodoModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DueDate     string `json:"due_date"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}
