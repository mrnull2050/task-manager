package storage

import (
	"encoding/json"
	"os"
	"task-manager/intenal/core"
)

type JsonStorage struct {
	Path string
}

func NewJsonStorage(path string) *JsonStorage {
	return &JsonStorage{Path: path}
}

func (s *JsonStorage) LoadTask() ([]core.Task, error) {
	if _, err := os.Stat(s.Path); os.IsNotExist(err) {
		return []core.Task{}, nil
	}
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, err
	}
	var tasks []core.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (s *JsonStorage) SaveTask(tasks []core.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0644)
}
