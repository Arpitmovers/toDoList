package main

import (
	"fmt"
	"os"

	"github.com/Arpitmovers/todolist"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [add|list|delete] 'task title'")
		return
	}

	command := os.Args[1]
	fmt.Println("command is ", command)

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add 'task title'")
			return
		}
		title := os.Args[2]
		todolist.AddTask(title)

	case "list":
		tasks, err := todolist.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		for _, task := range tasks {
			fmt.Printf("%d: %s [%t]\n", task.Id, task.Title, task.Completed)
		}

	default:
		fmt.Println("Unknown command passed")
	}
}
