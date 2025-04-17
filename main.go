package main

import (
	"fmt"
	"os"
	"time"
	"todo-cli/commands"
	"todo-cli/db"
	"todo-cli/models"

	"github.com/robfig/cron/v3"
)

func main() {
	db.Init() // Initialize DB on startup

	// Set up a cron job to notify daily at 9 AM
	c := cron.New()
	c.AddFunc("0 9 * * *", notifyPendingTasks) // Set to notify at 9 AM daily
	c.Start()

	// Handle command line arguments
	args := os.Args[1:]
	if len(args) < 1 {
		commands.ShowHelp()
		return
	}

	cmd := args[0]
	switch cmd {
	case "add":
		commands.Add(args[1:])
	case "list":
		commands.List()
	case "done":
		commands.Complete(args[1:])
	case "delete":
		commands.Delete(args[1:])
	default:
		commands.ShowHelp()
	}
}

// notifyPendingTasks will notify you of pending tasks every morning at 9 AM
func notifyPendingTasks() {
	tasks := models.GetAllTasks()

	var todaysTasks []models.Task
	var otherTasks []models.Task

	// Filter tasks by due date
	for _, t := range tasks {
		if t.DueDate == "" {
			otherTasks = append(otherTasks, t)
		} else if t.DueDate == time.Now().Format("2006-01-02") && !t.Done {
			todaysTasks = append(todaysTasks, t)
		} else {
			otherTasks = append(otherTasks, t)
		}
	}

	// Notify for today's tasks first
	fmt.Println("Today's Pending Tasks:")
	for _, task := range todaysTasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s (Priority: %s)\n", task.ID, status, task.Description, task.Priority)
	}

	// Notify for other tasks after
	fmt.Println("\nOther Pending Tasks:")
	for _, task := range otherTasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s (Priority: %s)\n", task.ID, status, task.Description, task.Priority)
	}
}
