package core

type Storage interface {
	LoadTask() ([]Task, error)
	SaveTask([]Task) error
}
