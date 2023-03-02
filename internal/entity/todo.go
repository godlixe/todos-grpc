package entity

import "time"

// Todo is an entity of Todo object
type Todo struct {
	ID          int64
	Title       string
	Description string
	IsDone      bool
	UserID      int64
	TimeStamp   time.Time
}
