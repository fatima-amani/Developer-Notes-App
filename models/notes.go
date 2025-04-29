package models

import "time"

type Note struct {
	ID        string    `json:"id"`         // UUID
	UserID    string    `json:"user_id"`    // Reference to user
	Title     string    `json:"title"`      // Note title
	Content   string    `json:"content"`    // Note body
	Tags      []string  `json:"tags"`       // Slice of tags
	CreatedAt time.Time `json:"created_at"` // Timestamp
	UpdatedAt time.Time `json:"updated_at"` // Timestamp
}
