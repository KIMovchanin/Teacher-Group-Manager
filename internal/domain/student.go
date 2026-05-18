package domain

import "time"

type Student struct {
	ID        int64
	FirstName string
	LastName  string
	CreatedAt time.Time
}
