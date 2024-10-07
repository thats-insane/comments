package data

import "time"

type Comment struct {
	ID        int64
	Content   string
	Author    string
	CreatedAt time.Time
	Version   int32
}
