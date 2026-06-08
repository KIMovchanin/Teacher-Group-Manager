package domain

import "time"

type Enrollment struct {
	StudentID int64
	GroupID   int64
	CreatedAt time.Time
}
