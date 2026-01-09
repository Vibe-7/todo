package model

import "time"

type Task struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreateAt    time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt"`
}
