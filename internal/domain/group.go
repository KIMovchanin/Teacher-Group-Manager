package domain

import "time"

type Group struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}
