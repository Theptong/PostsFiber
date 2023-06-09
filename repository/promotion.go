package repository

import (
	"time"
)

type PromotionRepository interface {
	CheckCollectionById() (MockPosts, error)
	GetCollection([]MockPosts) ([]MockPosts, error)
	CreateCollection() (MockPosts, error)
	// UpdateCollection(id string) (title,content string,published bool, err error)
	// DeleteCollection(id string) ([]MockPosts, error)
}

type MockPosts struct {
	Id        string
	Title     string
	Content   string
	Published bool
	ViewCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}
