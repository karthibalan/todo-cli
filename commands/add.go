package commands

import (
	"fmt"
	"strings"
	"todo-cli/models"
)

func Add(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: todo add <description> [-due YYYY-MM-DD] [-p Low|Medium|High]")
		return
	}

	desc := args[0]
	due := ""
	priority := "Low"

	for i, arg := range args {
		if arg == "-due" && i+1 < len(args) {
			due = args[i+1]
		}
		if arg == "-p" && i+1 < len(args) {
			p := strings.Title(strings.ToLower(args[i+1]))
			if p == "Low" || p == "Medium" || p == "High" {
				priority = p
			}
		}
	}

	models.AddTask(desc, due, priority)
}
