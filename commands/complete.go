package commands

import (
	"fmt"
	"strconv"
	"todo-cli/models"
)

func Complete(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: todo done <task_id>")
		return
	}
	id, _ := strconv.Atoi(args[0])
	models.CompleteTask(id)
}
