// file storage handling

package todolist

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const fileName = "/home/arpit/code/personal/toDoList/tasks.json"

func LoadTasks() ([]Task, error) {

	data, err := os.ReadFile(fileName)
	var tasks []Task
	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("error file donest exist  ", fileName, err)
			return []Task{}, nil
		}

		fmt.Println("error in reading  ", fileName, err)
		return []Task{}, err
	}

	unmarshErr := json.Unmarshal(data, &tasks)
	if unmarshErr != nil {
		fmt.Println("error in unmarshalling  ", unmarshErr)
		return []Task{}, unmarshErr
	}
	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return err
	}

	writeErr := os.WriteFile(fileName, data, 0644)

	if writeErr != nil {
		fmt.Println("error in write", writeErr)
		return writeErr
	}
	return nil
}
