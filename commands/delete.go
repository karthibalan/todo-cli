package commands

import (
	"fmt"
	"strconv"
	"todo-cli/models"
)

func Delete(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: todo delete <task_id>")
		return
	}
	id, _ := strconv.Atoi(args[0])
	models.DeleteTask(id)
}
