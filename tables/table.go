package table

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/rishabdhar12/go-todo/models"
)

func PrintTable(list []models.TodoModel) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{

			{Text: "ID"},
			{Text: "Name"},
			{Text: "DueDate"},
			{Text: "Status"},
			{Text: "Description"},
		},
	}

	var cells [][]*simpletable.Cell
	var status string

	for _, row := range list {
		id := row.ID
		name := row.Name
		due_date := row.DueDate
		if row.Status == true {
			status = "Complete"
		} else {
			status = "Pending"
		}
		description := row.Description

		cells = append(cells, []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: id},
			{Align: simpletable.AlignLeft, Text: name},
			{Align: simpletable.AlignLeft, Text: due_date},
			{Align: simpletable.AlignLeft, Text: status},
			{Align: simpletable.AlignLeft, Text: description},
		})

		table.Body = &simpletable.Body{Cells: cells}
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}
