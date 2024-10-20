package domain

import "time"

type Announcement struct {
	ChatID    string
	ID        int
	CreateBy  int
	Content   string
	MessageID *int
	CreateAt  time.Time
	DeletedAt *time.Time
}
