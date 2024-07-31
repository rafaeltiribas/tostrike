package model

import "time"

type Task struct {
	ID          int       `json:"id_task"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	IsCompleted bool      `json:"is_completed"`
}
