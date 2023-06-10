package service

import "workshop/structs"

type CollectionService interface {
	GetCollection() ([]structs.Posts, error)
	GetCollectionService() (structs.ListPosts)
	GetCollectionServiceById(id string) (*structs.Posts, error)
	GetCollectionServiceByListId(id string) ([]structs.Posts, error)
	CreateNewCollection(title, content string, published bool) (*structs.Posts, error)
	UpdateCollection(id, title, content string, published bool) (*structs.Posts, error)
	GetCollectionServiceLimit(page,limit int) structs.ListPosts
	DeleteCollection(id string) error
	GetServiceLimit(page,limit int) ([]structs.Posts, error)
}

