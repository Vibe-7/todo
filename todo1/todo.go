package todo1

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool
	CreateAt    time.Time
	CompletedAt *time.Time
}

func NewTasks(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,

		CreateAt:    time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	copleteTime := time.Now()

	t.Completed = true
	t.CompletedAt = &copleteTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompletedAt = nil
}
