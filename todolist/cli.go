package todolist

import (
	"fmt"
	"time"
)

func AddTask(title string) {
	tasks, err := LoadTasks()

	if err != nil {
		fmt.Println("err in load task ", err)
	}
	newTask := Task{Id: len(tasks) + 1, Title: title, CreatedAt: time.Now()}
	tasks = append(tasks, newTask)

	SaveTasks(tasks)
	fmt.Println("Task added:", title, tasks)
}

// func ListTasks() {
// 	tasks, _ := LoadTasks()
// 	if len(tasks) == 0 {
// 		fmt.Println("No tasks found.")
// 		return
// 	}
// 	for _, task := range tasks {
// 		status := "❌"
// 		if task.Completed {
// 			status = "✅"
// 		}
// 		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
// 	}
// }

// func CompleteTask(id int) {
// 	tasks, _ := LoadTasks()
// 	for i, task := range tasks {
// 		if task.ID == id {
// 			tasks[i].Completed = true
// 			SaveTasks(tasks)
// 			fmt.Println("Task marked as completed:", task.Title)
// 			return
// 		}
// 	}
// 	fmt.Println("Task not found.")
// }

// func DeleteTask(id int) {
// 	tasks, _ := LoadTasks()
// 	newTasks := []Task{}
// 	for _, task := range tasks {
// 		if task.ID != id {
// 			newTasks = append(newTasks, task)
// 		}
// 	}
// 	SaveTasks(newTasks)
// 	fmt.Println("Task deleted.")
// }
