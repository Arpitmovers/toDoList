package todolist

import "time"

type Task struct {
	id        int       `json:""id`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_At"`
}
