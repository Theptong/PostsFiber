package structs

import (
	"time"

	"github.com/google/uuid"
)

type ListPosts struct {
	Posts     []Posts `json:"posts"`
	Count     int   `json:"count"` 
	Limit     int   `json:"limit"` 
	Page      int   `json:"page"`
	TotalPage int   `json:"total_page"`
}

type Posts struct {
	Id        *uuid.UUID `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Published bool      `json:"published,omitempty"`
	ViewCount *int       `json:"view_count,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type PostsDB struct {
	Id        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Published bool      `db:"published"`
	ViewCount int       `db:"view_count"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

