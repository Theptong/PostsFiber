package repository

import "workshop/structs"

type CollectionRepository interface {
	GetAll() ([]structs.PostsDB, error)
	GetById(id string) (*structs.PostsDB, error)
	GetByTitle(id string) (*structs.PostsDB, error)
	GetByContent(id string) (*structs.PostsDB, error)
	GetByPublished(id string) ([]structs.PostsDB, error)
	GetByDate(id string,today string) ([]structs.PostsDB, error)
	CreateNewCollection(title, content string, published bool) (*structs.PostsDB, error)
	UpdateCollection(id,title, content string, published bool) (*structs.PostsDB, error)
	DeleteCollection(id string) (*structs.PostsDB, error)
	LimitCollection(page int,limit int) ([]structs.PostsDB, error)
}
