package models

import (
	"time"
)

const (
	TaskStatusTodo       TaskStatus = 0
	TaskStatusInProgress TaskStatus = 1
	TaskStatusDone       TaskStatus = 2
)

type (
	TaskStatus uint8

	Task struct {
		ID          uint64     `db:"id"          json:"id"`
		UserID      uint64     `db:"user_id"     json:"userId"`
		Title       string     `db:"title"       json:"title"`
		Description string     `db:"description" json:"description"`
		DueDate     time.Time  `db:"due_date"    json:"dueDate"`
		Status      TaskStatus `db:"status"      json:"status"`
		CreatedAt   time.Time  `db:"created_at"  json:"createdAt"`
		UpdatedAt   time.Time  `db:"updated_at"  json:"updatedAt"`
		DeletedAt   time.Time  `db:"deleted_at"  json:"deletedAt"`
	}
)

func (ts TaskStatus) String() string {
	switch ts {
	case TaskStatusTodo:
		return "todo"
	case TaskStatusInProgress:
		return "in_progress"
	case TaskStatusDone:
		return "done"
	default:
		return "unknown"
	}
}
