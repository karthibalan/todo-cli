package commands

import (
	"fmt"
	"todo-cli/models"
)

func List() {
	tasks := models.GetAllTasks()
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s (Due: %s, Priority: %s)\n", t.ID, status, t.Description, t.DueDate, t.Priority)
	}
}
