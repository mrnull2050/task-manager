package core

import (
	"fmt"
	"time"
)

type TaskService struct {
	storage Storage
}

func NewTaskService(s Storage) *TaskService {
	return &TaskService{storage: s}
}

func (ts *TaskService) Add(title string) error {
	tasks, err := ts.storage.LoadTask()
	if err != nil {
		return err
	}
	NewId := 1
	if len(tasks) > 0 {
		NewId = tasks[len(tasks)-1].ID + 1
	}
	task := Task{
		ID:        NewId,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now().Unix(),
	}
	tasks = append(tasks, task)
	return ts.storage.SaveTask(tasks)
}

func (ts *TaskService) List() ([]Task, error) {
	return ts.storage.LoadTask()
}

func (ts *TaskService) Done(id int) error {
	tasks, err := ts.storage.LoadTask()
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return ts.storage.SaveTask(tasks)
		}
	}
	return fmt.Errorf("task %d not found", id)
}
func (ts *TaskService) Delete(id int) error {
	tasks, err := ts.storage.LoadTask()
	if err != nil {
		return err
	}
	newList := make([]Task, 0)
	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newList = append(newList, t)
	}
	if !found {
		return fmt.Errorf("task %d not found", id)
	}
	return ts.storage.SaveTask(newList)
}
func (ts *TaskService) ClearDone() error {
	tasks, err := ts.storage.LoadTask()
	if err != nil {
		return err
	}
	newList := make([]Task, 0)
	for _, t := range tasks {
		if !t.Done {
			newList = append(newList, t)
		}
	}
	return ts.storage.SaveTask(newList)
}
