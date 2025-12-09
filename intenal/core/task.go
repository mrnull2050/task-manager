package core

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt int64  `json:"created_at"`
}
