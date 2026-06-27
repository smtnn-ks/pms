package models

import "time"

type TaskChangelog struct {
	ID uint

	TaskID uint
	Task   Task

	CreatedAt time.Time
}
