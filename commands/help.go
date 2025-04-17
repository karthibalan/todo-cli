package commands

import "fmt"

func ShowHelp() {
	fmt.Println("To-Do CLI App")
	fmt.Println("Usage:")
	fmt.Println("  todo add \"Task description\" -due 2025-04-20 -p High")
	fmt.Println("  todo list")
	fmt.Println("  todo done <task_id>")
	fmt.Println("  todo delete <task_id>")
}
