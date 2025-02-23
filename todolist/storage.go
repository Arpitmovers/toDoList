// file storage handling

package todolist

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func LoadTasks() ([]Task, error) {

	data, err := os.ReadFile(fileName)
	var tasks []Task
	if err != nil {
		return []Task{}, err
	}

	unmarshErr := json.Unmarshal(data, &tasks)
	if unmarshErr != nil {
		return []Task{}, unmarshErr
	}
	return tasks, nil
}

func SaveTask(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "")

	if err != nil {
		return err
	}

	writeErr := os.WriteFile(fileName, data, 777)
	if writeErr != nil {
		return writeErr
	}
}
