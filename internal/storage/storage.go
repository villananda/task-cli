package storage

import (
	"encoding/json"
	"io"
	"os"

	"task-cli/internal/task"
)

type Storage struct {
	FileName string
}

func NewStorage(fileName string) *Storage {
	return &Storage{FileName: fileName}
}

func (s *Storage) Load() ([]task.Task, error) {
	file, err := os.Open(s.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(bytes) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task
	if err := json.Unmarshal(bytes, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Storage) Save(tasks []task.Task) error {
	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, bytes, 0644)
}
