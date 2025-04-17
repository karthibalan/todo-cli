package models

import (
	"fmt"
	"log"
	"todo-cli/db"
)

type Task struct {
	ID          int
	Description string
	DueDate     string
	Priority    string
	Done        bool
}

func AddTask(desc, due, priority string) {
	_, err := db.DB.Exec("INSERT INTO tasks (description, due_date, priority) VALUES (?, ?, ?)", desc, due, priority)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task added.")
}

func GetAllTasks() []Task {
	rows, err := db.DB.Query("SELECT id, description, due_date, priority, done FROM tasks ORDER BY priority DESC, due_date")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Description, &t.DueDate, &t.Priority, &t.Done)
		tasks = append(tasks, t)
	}
	return tasks
}

func CompleteTask(id int) {
	_, err := db.DB.Exec("UPDATE tasks SET done = 1 WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task marked as done.")
}

func DeleteTask(id int) {
	_, err := db.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task deleted.")
}
