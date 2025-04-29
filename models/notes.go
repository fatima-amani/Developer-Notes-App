package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Note struct {
	ID        string         `gorm:"type:uuid;primaryKey" json:"id"`
	Username  string         `gorm:"index;not null" json:"user_name"`
	Title     string         `gorm:"not null" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"` // Use pq.StringArray for Postgres array
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Optional soft delete
}
